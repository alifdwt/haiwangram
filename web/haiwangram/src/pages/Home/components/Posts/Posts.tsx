import { Box, Heading } from "@chakra-ui/react";

export default function Posts() {
  return (
    <Box bg={"white"} _dark={{ bg: "gray.700" }} p={4} borderRadius={"lg"}>
      <Heading as="h1" size="2xl">
        Posts
      </Heading>
    </Box>
  );
}
