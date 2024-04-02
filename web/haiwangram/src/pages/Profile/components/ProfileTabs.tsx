import Post from "@/interface/Post";
import {
  Button,
  Flex,
  Image,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  Text,
} from "@chakra-ui/react";

type ProfileTabProps = {
  posts: Post[];
  postsNextPageButton: () => void;
};

export default function ProfileTabs({
  posts,
  postsNextPageButton,
}: ProfileTabProps) {
  return (
    <Tabs isFitted>
      <TabList mb={"1em"}>
        <Tab>Posts</Tab>
        <Tab>Likes</Tab>
      </TabList>
      <TabPanels>
        <TabPanel>
          <PostTab posts={posts} postsNextPageButton={postsNextPageButton} />
        </TabPanel>
        <TabPanel>Like</TabPanel>
      </TabPanels>
    </Tabs>
  );
}

function PostTab({
  posts,
  postsNextPageButton,
}: {
  posts: Post[];
  postsNextPageButton: () => void;
}) {
  return (
    <>
      <Flex gap={2} flexWrap={"wrap"}>
        {posts.map((post) => {
          return (
            <PostContainer key={post.id}>
              <PostBody
                title={post.title}
                caption={post.caption}
                photo_url={post.photo_url}
              />
            </PostContainer>
          );
        })}
      </Flex>
      <Button onClick={postsNextPageButton} w={"full"}>
        Load More
      </Button>
    </>
  );
}

function PostContainer({
  children,
}: // ref,
{
  children: React.ReactNode;
  // ref?: LegacyRef<HTMLDivElement>;
}) {
  return (
    <Flex
      bg={"white"}
      p={4}
      rounded={"xl"}
      flexDir={"column"}
      gap={4}
      _dark={{ bg: "gray.700" }}
      maxW={"22vw"}
      // ref={ref}
    >
      {children}
    </Flex>
  );
}

function PostBody({
  title,
  caption,
  photo_url,
}: {
  title: string;
  caption: string;
  photo_url: string;
}) {
  return (
    <Flex flexDir={"column"} gap={4}>
      {/* <Box> */}
      <Image
        src={photo_url}
        h={"300px"}
        w={"300px"}
        objectFit={"cover"}
        rounded={"xl"}
      />
      {/* </Box> */}
      <Text fontWeight={"bold"} textTransform={"capitalize"}>
        {title}
      </Text>
      <Text fontSize={"sm"}>{caption}</Text>
    </Flex>
  );
}
