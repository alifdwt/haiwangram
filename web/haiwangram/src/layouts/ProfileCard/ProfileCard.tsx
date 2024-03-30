import { Avatar, Box, Flex, Text } from "@chakra-ui/react";

const userData = {
  followers: 75,
  following: 100,
  posts: 50,
};

export default function ProfileCard() {
  return (
    <Box
      bg={"white"}
      _dark={{ bg: "gray.700" }}
      borderRadius={"lg"}
      p={4}
      mb={4}
    >
      <Box bg={"gray.100"} _dark={{ bg: "gray.600" }} borderRadius={"lg"} p={4}>
        <Flex alignItems={"center"}>
          <Avatar
            name="Dan Abrahmov"
            src="https://randomuser.me/api/portraits/men/59.jpg"
            size={"md"}
          />
          <Box ml={4}>
            <Text fontWeight={"bold"} fontSize={"lg"}>
              Dan Abrahmov
            </Text>
            <Text color={"gray.500"} _dark={{ color: "white" }}>
              @dan_abrahmov
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
