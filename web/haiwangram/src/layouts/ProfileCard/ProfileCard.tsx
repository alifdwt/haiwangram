import { RootState } from "@/store";
import { Avatar, Box, Flex, Link, Text } from "@chakra-ui/react";
import { useSelector } from "react-redux";

const userData = {
  followers: 75,
  following: 100,
  posts: 50,
};

export default function ProfileCard() {
  const { user } = useSelector((state: RootState) => state.user);

  return (
    <Box
      bg={"white"}
      _dark={{ bg: "gray.700" }}
      borderRadius={"lg"}
      p={4}
      mb={4}
    >
      <Box bg={"gray.100"} _dark={{ bg: "gray.600" }} borderRadius={"lg"} p={4}>
        <Flex alignItems={"center"} as={Link} href={`/${user?.username}`}>
          <Avatar
            name={user?.full_name}
            src={user?.profile_image_url}
            size={"md"}
          />
          <Box ml={4}>
            <Text fontWeight={"bold"} fontSize={"lg"}>
              {user?.full_name}
            </Text>
            <Text color={"gray.500"} _dark={{ color: "white" }}>
              @{user?.username}
            </Text>
          </Box>
        </Flex>

        <Flex mt={4} justify={"space-around"}>
          <Box textAlign={"center"}>
            <Text fontWeight={"bold"} fontSize={"lg"}>
              {userData.followers}
            </Text>
            <Text>Followers</Text>
          </Box>
          <Box textAlign={"center"}>
            <Text fontWeight={"bold"} fontSize={"lg"}>
              {userData.following}
            </Text>
            <Text>Following</Text>
          </Box>
          <Box textAlign={"center"}>
            <Text fontWeight={"bold"} fontSize={"lg"}>
              {userData.posts}
            </Text>
            <Text>Posts</Text>
          </Box>
        </Flex>
      </Box>
    </Box>
  );
}
