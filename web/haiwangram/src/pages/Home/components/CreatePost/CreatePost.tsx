import {
  Avatar,
  Box,
  Button,
  Flex,
  Image,
  Input,
  Select,
  Textarea,
} from "@chakra-ui/react";
import useCreatePost from "./hooks/useCreatePost";
import { HashIcon, ImagePlusIcon, Paperclip, X } from "lucide-react";

export default function CreatePost() {
  const {
    title,
    caption,
    image,
    setImage,
    user,
    handleTitleChange,
    handleCaptionChange,
    handleSubmit,
    handleFileUpload,
  } = useCreatePost();

  return (
    <Flex
      bg={"white"}
      _dark={{ bg: "gray.700" }}
      p={4}
      borderRadius={"lg"}
      gap={4}
      flexDir={"column"}
    >
      <Box as={Flex} align={"center"} gap={4}>
        <Box h={"full"}>
          <Avatar name={user?.full_name} src={user?.profile_image_url} />
        </Box>
        <Flex flexDir={"column"} gap={2} flex={5}>
          <Input
            value={title}
            onChange={handleTitleChange}
            placeholder="Judul"
            _focusVisible={{ border: "none" }}
          />
          <Textarea
            value={caption}
            onChange={handleCaptionChange}
            placeholder="Caption"
            _focusVisible={{ border: "none" }}
            resize={"none"}
          />
        </Flex>
        <Button
          h={"full"}
          onClick={handleSubmit}
          bg={"primary.200"}
          color={"primary.700"}
          _dark={{ bg: "primary.700", color: "white" }}
          _hover={{ bg: "primary.400" }}
        >
          Post
        </Button>
      </Box>
      {/* <Divider /> */}
      <Flex justifyContent={"space-between"} alignItems={"center"}>
        <Flex flex={4}>
          <Button
            leftIcon={<ImagePlusIcon color="#1b79e5" size={20} />}
            onClick={handleFileUpload}
            variant={"ghost"}
            color={"gray.500"}
            px={4}
            py={2}
          >
            Image
          </Button>
          <Button
            leftIcon={<Paperclip color="orange" size={20} />}
            variant={"ghost"}
            color={"gray.500"}
            px={4}
            py={2}
          >
            Attachment
          </Button>
          <Button
            leftIcon={<HashIcon color="green" size={20} />}
            variant={"ghost"}
            color={"gray.500"}
            px={4}
            py={2}
          >
            Hashtag
          </Button>
        </Flex>
        {image && (
          <Box as={Flex} flex={1}>
            <Box w="fit-content" position="relative">
              <Button
                bg="transparent"
                _hover={{ bg: "transparent" }}
                position="absolute"
                top={0}
                right={0}
                onClick={() => setImage(null)}
              >
                <X color="white" />
              </Button>
              <Image
                src={URL.createObjectURL(image)}
                rounded={"lg"}
                w={"40px"}
              />
            </Box>
          </Box>
        )}
        <Select size={"sm"} flex={1} variant={"unstyled"} color={"gray.500"}>
          <option value="public">Public</option>
          <option value="private">Private</option>
        </Select>
      </Flex>
    </Flex>
  );
}
