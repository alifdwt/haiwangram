import { Box, Button, Flex, Link, Text } from "@chakra-ui/react";
import { ArrowLeftIcon } from "lucide-react";

export default function PostTopBar() {
  return (
    <Flex borderBottom={"1px"} pb={2} alignItems={"center"}>
      <Button variant={"ghost"} as={Link} href="/">
        <ArrowLeftIcon />
      </Button>
      <Box>
        <Text fontWeight={"bold"}>Post</Text>
      </Box>
    </Flex>
  );
}
