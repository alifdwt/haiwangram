import { Button } from "@chakra-ui/react";
import { HeartIcon } from "lucide-react";
import React from "react";
import useLikeButton from "./hooks/useLikeButton";

type LikeButtonProps = {
  likeCount: number;
  isLiked: boolean;
  postId: number;
};

const LikeButton: React.FC<LikeButtonProps> = ({
  likeCount,
  isLiked,
  postId,
}) => {
  const { handleLike } = useLikeButton({ isLiked, postId });
  return (
    <Button
      leftIcon={isLiked ? <HeartIcon fill="red" /> : <HeartIcon />}
      variant={"ghost"}
      color={isLiked ? "red" : "gray"}
      onClick={handleLike}
    >
      {likeCount} {likeCount === 1 ? "Like" : "Likes"}
    </Button>
  );
};

export default LikeButton;
