import axiosFetch from "@/config/axiosFetch";
import { getBookmarks } from "@/slices/bookmark/bookmarkSlice";
import { RootState } from "@/store";
import { Box, Button, Flex, Link, Skeleton, Text } from "@chakra-ui/react";
import { useQuery } from "@tanstack/react-query";
import { useDispatch, useSelector } from "react-redux";
import {
  PostBody,
  PostContainer,
  PostFooter,
  PostHeader,
} from "../Home/components/Posts/Posts";
import { ArrowLeftIcon } from "lucide-react";

export default function BookmarkPage() {
  const dispatch = useDispatch();
  const { bookmarks } = useSelector((state: RootState) => state.bookmark);
  const { accessToken } = useSelector((state: RootState) => state.user);

  const { isLoading } = useQuery({
    queryKey: ["bookmarks"],
    queryFn: async () => {
      const { data } = await axiosFetch.get("/bookmarks/", {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
      dispatch(getBookmarks(data));
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
      <Flex borderBottom={"1px"} pb={2} alignItems={"center"}>
        <Button variant={"ghost"} as={Link} href="/">
          <ArrowLeftIcon />
        </Button>
        <Box>
          <Text fontWeight={"bold"}>Bookmark</Text>
        </Box>
      </Flex>
      {bookmarks.map((bookmark) => {
        return (
          <PostContainer key={bookmark.id}>
            <PostHeader
              profile_image_url={
                bookmark.photo.user?.profile_image_url as string
              }
              full_name={bookmark.photo.user?.full_name as string}
              username={bookmark.photo.user?.username as string}
              created_at={bookmark.photo.created_at}
            />
            <PostBody
              title={bookmark.photo.title}
              caption={bookmark.photo.caption}
              photo_url={bookmark.photo.photo_url}
              post_id={bookmark.photo.id}
            />
            <PostFooter
              postId={bookmark.photo.id}
              isLiked={false}
              isBookmarked={true}
              likeCount={0}
              post={bookmark.photo}
            />
          </PostContainer>
        );
      })}
    </>
  );
}
