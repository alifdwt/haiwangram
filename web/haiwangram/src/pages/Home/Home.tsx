import { Divider, Flex } from "@chakra-ui/react";
import HaiwanStories from "./components/HaiwanStories/HaiwanStories";
import CreatePost from "./components/CreatePost/CreatePost";
import Posts from "./components/Posts/Posts";

export default function Home() {
  return (
    <Flex
      flexDir={"column"}
      gap={4}
      css={{ "&::-webkit-scrollbar": { display: "none" } }}
    >
      <HaiwanStories />
      <CreatePost />
      <Divider borderWidth={"2px"} />
      <Posts />
    </Flex>
  );
}
