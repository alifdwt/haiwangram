import { Box, Heading } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import PostTopBar from "./components/PostTopBar";

export default function Post() {
  const { postId } = useParams();

  return (
    <Box>
      <PostTopBar />
      <Heading color={"primary.700"}>{postId}</Heading>
    </Box>
  );
}
