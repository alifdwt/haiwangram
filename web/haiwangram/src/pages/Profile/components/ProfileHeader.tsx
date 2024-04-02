import User from "@/interface/User";
import { Avatar, Button, Flex, Text } from "@chakra-ui/react";
import { EllipsisIcon } from "lucide-react";

type ProfileHeaderProps = {
  user: User;
};

export default function ProfileHeader({ user }: ProfileHeaderProps) {
  return (
    <Flex gap={5} p={4} alignItems={"center"}>
      <Avatar name={user.full_name} src={user.profile_image_url} size={"2xl"} />
      <Flex flexDir={"column"} gap={2}>
        <Flex gap={4} alignItems={"center"}>
          <Text fontWeight={"bold"} fontSize={"lg"}>
            {user.username}
          </Text>
          <Button w={"100px"}>Ikuti</Button>
          <Button w={"100px"}>Pesan</Button>
          <Button variant={"ghost"}>
            <EllipsisIcon size={20} />
          </Button>
        </Flex>
        <Flex gap={4}>
          <Text fontWeight={"bold"}>
            {123}{" "}
            <Text as={"span"} fontWeight={"normal"}>
              Post
            </Text>
          </Text>
          <Text fontWeight={"bold"}>
            {123}{" "}
            <Text as={"span"} fontWeight={"normal"}>
              Pengikut
            </Text>
          </Text>
          <Text fontWeight={"bold"}>
            {123}{" "}
            <Text as={"span"} fontWeight={"normal"}>
              Mengikuti
            </Text>
          </Text>
        </Flex>
        <Text>{user.full_name}</Text>
        <Text>{user.description}</Text>
      </Flex>
    </Flex>
  );
}
