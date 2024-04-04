import { Button } from "@chakra-ui/react";
import useBookmarkButton from "./hooks/useBookmarkButton";
import { BookmarkIcon } from "lucide-react";

type BookmarkButtonProps = {
  isBookmarked: boolean;
  postId: number;
};

export default function BookmarkButton({
  isBookmarked,
  postId,
}: BookmarkButtonProps) {
  const { handleBookmark } = useBookmarkButton({ isBookmarked, postId });

  return (
    <Button
      variant={"ghost"}
      color={isBookmarked ? "primary.700" : "gray"}
      onClick={handleBookmark}
    >
      {isBookmarked ? <BookmarkIcon fill="#1b79e5" /> : <BookmarkIcon />}
    </Button>
  );
}
