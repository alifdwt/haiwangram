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
  Link,
  Skeleton,
  Text,
} from "@chakra-ui/react";
import { useInfiniteQuery } from "@tanstack/react-query";
import { BookmarkIcon, EllipsisVerticalIcon } from "lucide-react";
import React, { LegacyRef } from "react";
import { useDispatch, useSelector } from "react-redux";
import LikeButton from "./components/LikeButton";
import CommentButton from "./components/CommentButton";
import Post from "@/interface/Post";

function PostContainer({
  children,
  ref,
}: {
  children: React.ReactNode;
  ref?: LegacyRef<HTMLDivElement>;
}) {
  return (
    <Flex
      bg={"white"}
      p={4}
      rounded={"xl"}
      flexDir={"column"}
      gap={4}
      _dark={{ bg: "gray.700" }}
      ref={ref}
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
          <Link
            href={`/${username}`}
            display={"flex"}
            gap={2}
            alignItems={"center"}
          >
            <Text>{full_name}</Text>
            <Text color={"gray"} fontSize={"sm"}>
              @{username}
            </Text>
          </Link>
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
  postId,
  likeCount,
  isLiked,
  post,
}: {
  postId: number;
  likeCount: number;
  isLiked: boolean;
  post: Post;
}) {
  return (
    <Flex justifyContent={"space-between"}>
      <Flex>
        <LikeButton likeCount={likeCount} isLiked={isLiked} postId={postId} />
        <CommentButton post={post} isLiked={isLiked} />
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
  const { user } = useSelector((state: RootState) => state.user);
  const userId = Number(user?.id);

  const { fetchNextPage, isFetchingNextPage, isLoading } =
    // @ts-expect-error next-line
    useInfiniteQuery({
      queryKey: ["posts"],
      queryFn: async ({ pageParam = 1 }) => {
        const { data } = await axiosFetch.get(
          `/photos/?limit=${pageParam * 4}`
        );
        dispatch(getPosts(data));
        return data;
      },
      getNextPageParam: (_, pages) => {
        return pages.length + 1;
      },
      initialData: {
        pages: [posts],
        pageParams: [1],
      },
    });

  // const lastPostRef = useRef<HTMLElement>(null);
  // const { ref, entry } = useIntersection({
  //   root: lastPostRef.current,
  //   threshold: 1,
  // });

  // useEffect(() => {
  //   if (entry?.isIntersecting) fetchNextPage();
  // }, [entry, fetchNextPage]);

  // const _posts = data?.pages.flatMap((page) => page);

  if (isLoading) {
    return (
      <Box>
        <Skeleton height={"300px"} />
      </Box>
    );
  }

  return (
    <>
      {posts.map((post) => {
        const isLiked = post.likes?.some((like) => like.user_id === userId);
        // if (i === _posts.length - 1)
        //   return (
        //     <PostContainer key={post.id} ref={ref}>
        //       <PostHeader
        //         profile_image_url={post.user?.profile_image_url as string}
        //         full_name={post.user?.full_name as string}
        //         username={post.user?.username as string}
        //         created_at={post.created_at as string}
        //       />
        //       <PostBody
        //         title={post.title}
        //         caption={post.caption}
        //         photo_url={post.photo_url}
        //       />
        //       <PostFooter
        //         postId={post.id}
        //         isLiked={isLiked!}
        //         likeCount={post.likes?.length as number}
        //         commentCount={post.comments?.length as number}
        //       />
        //     </PostContainer>
        //   );

        return (
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
              postId={post.id}
              isLiked={isLiked!}
              likeCount={post.likes?.length as number}
              post={post}
            />
          </PostContainer>
        );
      })}
      <Button
        onClick={() => fetchNextPage()}
        isLoading={isFetchingNextPage}
        loadingText="Memuat..."
      >
        Load More
      </Button>
    </>
  );
}
