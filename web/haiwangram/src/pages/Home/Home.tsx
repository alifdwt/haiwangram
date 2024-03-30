import { Flex } from "@chakra-ui/react";
import HaiwanStories from "./components/HaiwanStories/HaiwanStories";
import CreatePost from "./components/CreatePost/CreatePost";
import Posts from "./components/Posts/Posts";

export default function Home() {
  return (
    <Flex overflowY={"scroll"} flexDir={"column"} gap={4}>
      <HaiwanStories />
      <CreatePost />
      <Posts />
    </Flex>
  );
}
