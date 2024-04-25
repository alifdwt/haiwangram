import useFetchWithId from "@/hooks/useFetchWithId";
import useToast from "@/hooks/useToast";
import * as PostType from "@/interface/Post";
import BookmarkButton from "@/pages/Home/components/Posts/components/BookmarkButton";
import LikeButton from "@/pages/Home/components/Posts/components/LikeButton";
import { TimeConverter } from "@/util/converter";
import {
  Avatar,
  Box,
  Button,
  Flex,
  Image,
  Link,
  Spinner,
  Text,
} from "@chakra-ui/react";
import { ReactNode } from "react";
import { useNavigate } from "react-router-dom";

function PostContainer({ children }: { children: ReactNode }) {
  return (
    <Flex gap={4} py={4} flexDir={"column"}>
      {children}
    </Flex>
  );
}

function PostHeader({
  fullName,
  username,
  profileUrl,
}: {
  fullName: string;
  username: string;
  profileUrl: string;
}) {
  return (
    <Flex
      gap={2}
      as={Link}
      href={`/${username}`}
      alignItems={"center"}
      _hover={{ textDecoration: "none" }}
    >
      <Avatar name={fullName} src={profileUrl} />
      <Box>
        <Text fontWeight={"bold"} _hover={{ textDecoration: "underline" }}>
          {fullName}
        </Text>
        <Text color={"gray"}>@{username}</Text>
      </Box>
      {/* {isUser && } */}
    </Flex>
  );
}

function PostBody({
  title,
  caption,
  photo_url,
  created_at,
  updated_at,
}: {
  title: string;
  caption: string;
  photo_url: string;
  created_at: string;
  updated_at: string;
}) {
  return (
    <Flex flexDir={"column"} gap={4} _hover={{ textDecoration: "none" }}>
      {/* <Box> */}
      <Image
        src={photo_url}
        w={"full"}
        maxH={"70vh"}
        objectFit={"cover"}
        rounded={"xl"}
        alt={title}
      />
      {/* </Box> */}
      <Text fontWeight={"bold"} textTransform={"capitalize"}>
        {title}
      </Text>
      <Text fontSize={"sm"}>{caption}</Text>
      <Text color={"gray"}>
        {created_at === updated_at
          ? TimeConverter(created_at)
          : `Diperbarui ${TimeConverter(updated_at)}`}
      </Text>
    </Flex>
  );
}

function PostFooter({
  postId,
  likeCount,
  commentCount,
  isLiked,
  isBookmarked,
}: {
  postId: number;
  likeCount: number;
  commentCount: number;
  isLiked: boolean;
  isBookmarked: boolean;
}) {
  return (
    <Flex justifyContent={"space-between"}>
      <Flex>
        <LikeButton likeCount={likeCount} isLiked={isLiked} postId={postId} />
        <Button variant={"ghost"} color={"gray"}>
          {commentCount} Komentar
        </Button>
      </Flex>
      <BookmarkButton isBookmarked={isBookmarked} postId={postId} />
    </Flex>
  );
}

export default function Post({
  postId,
  userId,
}: {
  postId: string;
  userId: number;
}) {
  const navigate = useNavigate();
  const toast = useToast();
  const { data: postData, isLoading: isLoadingPostData } = useFetchWithId({
    queryKey: "post",
    id: postId,
    fetchRoute: "/photos",
  });

  if (isLoadingPostData) {
    return (
      <Flex align={"center"} justify={"center"} h={"full"}>
        <Spinner size={"lg"} />
      </Flex>
    );
  }

  const post: PostType.default = postData;
  if (post === undefined) {
    toast("Error", "Post tidak ditemukan", "error");
    return navigate("/");
  }

  const isLiked = post.likes?.some((like) => like.user_id === userId);
  const totalLikes = post.likes?.length;
  const totalComments = post.comments?.length;
  // const isUser = post.user_id === userId
  const isBookmarked = post.bookmarks?.some(
    (bookmark) => bookmark.user_id === userId
  );

  return (
    <PostContainer>
      <PostHeader
        fullName={post.user?.full_name as string}
        username={post.user?.username as string}
        profileUrl={post.user?.profile_image_url as string}
      />
      <PostBody
        title={post.title}
        caption={post.caption}
        photo_url={post.photo_url}
        created_at={post.created_at}
        updated_at={post.updated_at}
      />
      <PostFooter
        postId={parseInt(postId)}
        likeCount={totalLikes as number}
        commentCount={totalComments as number}
        isLiked={isLiked as boolean}
        isBookmarked={isBookmarked as boolean}
      />
    </PostContainer>
  );
}
