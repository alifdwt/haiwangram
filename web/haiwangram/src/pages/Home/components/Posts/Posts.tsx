import axiosFetch from "@/config/axiosFetch";
import { getPosts } from "@/slices/posts/postSlice";
import { RootState } from "@/store";
import { TimeConverter } from "@/util/converter";
import {
  Avatar,
  Box,
  Button,
  Flex,
  Image,
  Skeleton,
  Text,
} from "@chakra-ui/react";
import { useQuery } from "@tanstack/react-query";
import {
  BookmarkIcon,
  EllipsisVerticalIcon,
  HeartIcon,
  MessageCircleIcon,
} from "lucide-react";
import React from "react";
import { useDispatch, useSelector } from "react-redux";

// const posts = [
//   {
//     id: 1,
//     caption:
//       "lorem ipsum dolor sit amet consectetur adipiscing elit congue rhoncus nisi vulputate dignissim morbi luctus",
//     title: "lorem ipsum dolor sit amet consectetur adipiscing",
//     photo_url: "https://picsum.photos/id/10/1280/720",
//     user_id: 3,
//     user: {
//       id: 3,
//       email: "hyfkvm-cuzxyq@example.com",
//       username: "hyfkvm-cuzxyq",
//       full_name: "Hyfkvm Cuzxyq",
//       profile_image_url: "https://randomuser.me/api/portraits/women/48.jpg",
//       description: "lorem ipsum dolor sit amet consectetur adipiscing",
//     },
//     comments: [
//       {
//         id: 1,
//         message: "lorem ipsum dolor sit amet consectetur adipiscing",
//         photo_id: 1,
//         user_id: 3,
//       },
//     ],
//   },
//   {
//     id: 3,
//     caption:
//       "lorem ipsum dolor sit amet consectetur adipiscing elit congue rhoncus nisi vulputate dignissim morbi luctus",
//     title: "lorem ipsum dolor sit amet consectetur adipiscing",
//     photo_url: "https://randomuser.me/api/portraits/women/89.jpg",
//     user_id: 4,
//     user: {
//       id: 4,
//       email: "mlzgtc-klwzuh@example.com",
//       username: "mlzgtc-klwzuh",
//       full_name: "Mlzgtc Klwzuh",
//       profile_image_url: "https://randomuser.me/api/portraits/women/33.jpg",
//       description: "lorem ipsum dolor sit amet consectetur adipiscing",
//     },
//     comments: [
//       {
//         id: 3,
//         message: "lorem ipsum dolor sit amet consectetur adipiscing",
//         photo_id: 3,
//         user_id: 4,
//       },
//     ],
//   },
//   {
//     id: 5,
//     caption:
//       "lorem ipsum dolor sit amet consectetur adipiscing elit congue rhoncus nisi vulputate dignissim morbi luctus",
//     title: "lorem ipsum dolor sit amet consectetur adipiscing",
//     photo_url: "https://randomuser.me/api/portraits/women/33.jpg",
//     user_id: 5,
//     user: {
//       id: 5,
//       email: "mdcmtb-vzzmuf@example.com",
//       username: "mdcmtb-vzzmuf",
//       full_name: "Mdcmtb Vzzmuf",
//       profile_image_url: "https://randomuser.me/api/portraits/men/37.jpg",
//       description: "lorem ipsum dolor sit amet consectetur adipiscing",
//     },
//     comments: [
//       {
//         id: 5,
//         message: "lorem ipsum dolor sit amet consectetur adipiscing",
//         photo_id: 5,
//         user_id: 5,
//       },
//     ],
//   },
// ];

function PostContainer({ children }: { children: React.ReactNode }) {
  return (
    <Flex
      bg={"white"}
      p={4}
      rounded={"xl"}
      flexDir={"column"}
      gap={4}
      _dark={{ bg: "gray.700" }}
    >
      {children}
    </Flex>
  );
}

function PostHeader({
  profile_image_url,
  full_name,
  username,
  created_at,
}: {
  profile_image_url: string;
  full_name: string;
  username: string;
  created_at: string;
}) {
  return (
    <Flex justifyContent={"space-between"}>
      <Flex gap={2}>
        <Avatar name={full_name} src={profile_image_url} />
        <Box>
          <Flex gap={2} alignItems={"center"}>
            <Text>{full_name}</Text>
            <Text color={"gray"} fontSize={"sm"}>
              @{username}
            </Text>
          </Flex>
          <Text color={"gray"}>{TimeConverter(created_at)}</Text>
        </Box>
      </Flex>
      <Button variant={"ghost"}>
        <EllipsisVerticalIcon />
      </Button>
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
        w={"full"}
        maxH={"300px"}
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

function PostFooter({
  likeCount,
  commentCount,
}: {
  likeCount: number;
  commentCount: number;
}) {
  return (
    <Flex justifyContent={"space-between"}>
      <Flex>
        <Button leftIcon={<HeartIcon />} variant={"ghost"} color={"gray"}>
          {likeCount} {likeCount === 1 ? "Like" : "Likes"}
        </Button>
        <Button
          leftIcon={<MessageCircleIcon />}
          variant={"ghost"}
          color={"gray"}
        >
          {commentCount} {commentCount === 1 ? "Comment" : "Comments"}
        </Button>
      </Flex>
      <Button variant={"ghost"} color={"gray"}>
        <BookmarkIcon />
      </Button>
    </Flex>
  );
}

export default function Posts() {
  const dispatch = useDispatch();
  const { posts } = useSelector((state: RootState) => state.post);

  const { isLoading } = useQuery({
    queryKey: ["posts"],
    queryFn: async () => {
      const { data } = await axiosFetch.get("/photos/");
      dispatch(getPosts(data));
      return data;
    },
  });

  if (isLoading) {
    return (
      <Box>
        <Skeleton height={"300px"} />
      </Box>
    );
  }

  return (
    <>
      {posts.map((post) => (
        <PostContainer key={post.id}>
          <PostHeader
            profile_image_url={post.user?.profile_image_url as string}
            full_name={post.user?.full_name as string}
            username={post.user?.username as string}
            created_at={post.created_at as string}
          />
          <PostBody
            title={post.title}
            caption={post.caption}
            photo_url={post.photo_url}
          />
          <PostFooter
            likeCount={post.likes?.length as number}
            commentCount={post.comments?.length as number}
          />
        </PostContainer>
      ))}
    </>
  );
}
