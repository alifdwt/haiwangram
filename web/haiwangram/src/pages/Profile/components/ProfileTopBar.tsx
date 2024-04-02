import { Box, Button, Flex, Text } from "@chakra-ui/react";
import { ArrowLeftIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";

export default function ProfileTopBar({
  fullName,
  postCount,
}: {
  fullName: string;
  postCount: number;
}) {
  const navigate = useNavigate();
  return (
    <Flex borderBottom={"1px"} pb={2}>
      <Button variant={"ghost"} onClick={() => navigate(-1)}>
        <ArrowLeftIcon />
      </Button>
      <Box>
        <Text fontWeight={"bold"} fontSize={"sm"}>
          {fullName}
        </Text>
        <Text color={"gray"} fontSize={"sm"}>
          {postCount} {postCount === 1 ? "Post" : "Posts"}
        </Text>
      </Box>
    </Flex>
  );
}
