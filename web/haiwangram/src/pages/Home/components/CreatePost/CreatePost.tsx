import { Box, Heading } from "@chakra-ui/react";

export default function CreatePost() {
  return (
    <Box bg={"white"} _dark={{ bg: "gray.700" }} p={4} borderRadius={"lg"}>
      <Heading as="h1" size="2xl">
        Create Post
      </Heading>
    </Box>
  );
}
