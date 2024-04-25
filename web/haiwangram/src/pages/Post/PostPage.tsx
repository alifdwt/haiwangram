import { Box } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import PostTopBar from "./components/PostTopBar";
import { useSelector } from "react-redux";
import { RootState } from "@/store";
import Post from "./components/Post";

export default function PostPage() {
  const { user } = useSelector((state: RootState) => state.user);
  const userId = user?.id;

  const { postId } = useParams();

  return (
    <Box>
      <PostTopBar />
      {/* @ts-expect-error next-line */}
      <Post userId={userId as number} postId={postId as string} />
    </Box>
  );
}
