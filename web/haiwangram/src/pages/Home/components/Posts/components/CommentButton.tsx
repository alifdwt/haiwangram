import Post from "@/interface/Post";
import { getIndonesianDuration } from "@/util/converter";
import {
  Avatar,
  Box,
  Button,
  ButtonGroup,
  Flex,
  Grid,
  Image,
  Input,
  InputGroup,
  InputRightElement,
  Link,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Skeleton,
  Text,
  useDisclosure,
} from "@chakra-ui/react";
import {
  BookmarkIcon,
  HeartIcon,
  MessageCircleIcon,
  Share2Icon,
} from "lucide-react";
import useLikeButton from "./hooks/useLikeButton";
import useCommentButton from "./hooks/useCommentButton";
import axiosFetch from "@/config/axiosFetch";
import { useQuery } from "@tanstack/react-query";
import Comment from "@/interface/Comment";
import useBookmarkButton from "./hooks/useBookmarkButton";

type CommentButtonProps = {
  post: Post;
  isLiked: boolean;
  isBookmarked: boolean;
};

export default function CommentButton({
  post,
  isLiked,
  isBookmarked,
}: CommentButtonProps) {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { handleLike } = useLikeButton({ isLiked, postId: post.id });
  const { handleBookmark } = useBookmarkButton({
    isBookmarked,
    postId: post.id,
  });
  const { comment, handleCommentChange, handleSubmit, mutation } =
    useCommentButton({ postId: post.id });

  const { data, isLoading } = useQuery({
    queryKey: ["comments", post.id],
    queryFn: async () => {
      const { data } = await axiosFetch.get(`/comments/?photoId=${post.id}`);
      return data;
    },
  });

  const comments: Comment[] = data;

  return (
    <>
      <Button
        onClick={onOpen}
        leftIcon={<MessageCircleIcon />}
        variant={"ghost"}
        color={"gray"}
      >
        {post.comments?.length}{" "}
        {post.comments?.length === 1 ? "Comment" : "Comments"}
      </Button>

      <Modal isOpen={isOpen} onClose={onClose} isCentered size={"5xl"}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader as={Flex} gap={4}>
            <Avatar
              name={post.user?.full_name}
              src={post.user?.profile_image_url}
              size={"sm"}
            />
            <Text>{post.title}</Text>
          </ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Grid templateColumns={"repeat(2, 1fr)"} gap={4} h={"80vh"}>
              <Box h={"full"}>
                <Image src={post.photo_url} />
              </Box>
              <Flex flexDir={"column"} justifyContent={"space-between"}>
                <Flex gap={2} borderBottom={"1px"} pb={2}>
                  <Avatar
                    name={post.user?.full_name}
                    src={post.user?.profile_image_url}
                    size={"sm"}
                  />
                  <Text>
                    <Link href={`/${post.user?.username}`} fontWeight={"bold"}>
                      {post.user?.full_name}
                    </Link>{" "}
                    {post.caption}
                  </Text>
                </Flex>
                {isLoading ? (
                  <Box>
                    <Skeleton height="20px" />
                  </Box>
                ) : (
                  <Flex
                    flexDir={"column"}
                    gap={2}
                    h={"full"}
                    overflowY={"auto"}
                    pt={2}
                  >
                    {comments.map((comment) => (
                      <Flex gap={2}>
                        <Avatar
                          name="Zainal Abidin"
                          size={"sm"}
                          src={comment.user?.profile_image_url}
                        />
                        <Box>
                          <Text>
                            <Link
                              href={`/${comment.user?.username}`}
                              fontWeight={"bold"}
                            >
                              {comment.user?.username}
                            </Link>{" "}
                            {comment.message}
                          </Text>
                          <Text as={"span"} fontSize={"sm"} color={"gray"}>
                            {getIndonesianDuration(comment.created_at)}
                          </Text>
                        </Box>
                      </Flex>
                    ))}
                  </Flex>
                )}
                <Box borderTop={"1px"}>
                  <Flex justifyContent={"space-between"}>
                    <ButtonGroup>
                      <Button
                        variant={"ghost"}
                        color={isLiked ? "red" : ""}
                        onClick={handleLike}
                      >
                        {isLiked ? <HeartIcon fill="red" /> : <HeartIcon />}
                      </Button>
                      <Button variant={"ghost"}>
                        <MessageCircleIcon />
                      </Button>
                      <Button variant={"ghost"}>
                        <Share2Icon />
                      </Button>
                    </ButtonGroup>
                    <Button
                      variant={"ghost"}
                      color={isBookmarked ? "primary.700" : ""}
                      onClick={handleBookmark}
                    >
                      {isBookmarked ? (
                        <BookmarkIcon fill="#1b79e5" />
                      ) : (
                        <BookmarkIcon />
                      )}
                    </Button>
                  </Flex>
                  <Box p={2}>
                    <Text fontWeight={"bold"}>
                      {post.likes?.length}{" "}
                      {post.likes?.length === 1 ? "Like" : "Likes"}
                    </Text>
                    <Text color="gray">
                      {getIndonesianDuration(post.created_at)}
                    </Text>
                  </Box>
                  <Flex>
                    <InputGroup>
                      <Input
                        type="text"
                        variant={"flushed"}
                        placeholder="Add a comment..."
                        value={comment}
                        onChange={handleCommentChange}
                      />
                      <InputRightElement w={"auto"}>
                        <Button
                          onClick={handleSubmit}
                          isLoading={mutation.isPending}
                        >
                          Komen
                        </Button>
                      </InputRightElement>
                    </InputGroup>
                  </Flex>
                </Box>
              </Flex>
            </Grid>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
}
