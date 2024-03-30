import {
  Avatar,
  AvatarBadge,
  Button,
  Flex,
  Input,
  InputGroup,
  InputLeftElement,
  Stack,
  Text,
} from "@chakra-ui/react";
import { SearchIcon, SquarePenIcon } from "lucide-react";

const users = [
  {
    id: 2,
    email: "hofasa-fhyeot@example.com",
    username: "hofasa-fhyeot",
    full_name: "Hofasa Fhyeot",
    profile_image_url: "https://randomuser.me/api/portraits/men/90.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 14,
    email: "zuvwjw-iacdug@example.com",
    username: "zuvwjw-iacdug",
    full_name: "Zuvwjw Iacdug",
    profile_image_url: "https://randomuser.me/api/portraits/men/16.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 7,
    email: "zjnwgv-qssfzw@example.com",
    username: "zjnwgv-qssfzw",
    full_name: "Zjnwgv Qssfzw",
    profile_image_url: "https://randomuser.me/api/portraits/women/81.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 6,
    email: "kfnodv-metoje@example.com",
    username: "kfnodv-metoje",
    full_name: "Kfnodv Metoje",
    profile_image_url: "https://randomuser.me/api/portraits/women/56.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 19,
    email: "bludjt-nfpton@example.com",
    username: "bludjt-nfpton",
    full_name: "Bludjt Nfpton",
    profile_image_url: "https://randomuser.me/api/portraits/men/70.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 4,
    email: "pzowcj-qhuigo@example.com",
    username: "pzowcj-qhuigo",
    full_name: "Pzowcj Qhuigo",
    profile_image_url: "https://randomuser.me/api/portraits/men/41.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 5,
    email: "tlsuwe-vhsogp@example.com",
    username: "tlsuwe-vhsogp",
    full_name: "Tlsuwe Vhsogp",
    profile_image_url: "https://randomuser.me/api/portraits/men/6.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 13,
    email: "yutzra-jfcyng@example.com",
    username: "yutzra-jfcyng",
    full_name: "Yutzra Jfcyng",
    profile_image_url: "https://randomuser.me/api/portraits/women/21.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 15,
    email: "nuuwoo-ekjowu@example.com",
    username: "nuuwoo-ekjowu",
    full_name: "Nuuwoo Ekjowu",
    profile_image_url: "https://randomuser.me/api/portraits/men/47.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
  {
    id: 16,
    email: "vovcsp-trmqcp@example.com",
    username: "vovcsp-trmqcp",
    full_name: "Vovcsp Trmqcp",
    profile_image_url: "https://randomuser.me/api/portraits/men/40.jpg",
    description: "lorem ipsum dolor sit amet consectetur adipiscing",
  },
];

export default function Messages() {
  return (
    <Stack
      bg={"white"}
      _dark={{ bg: "gray.700" }}
      p={4}
      borderRadius={"xl"}
      gap={4}
    >
      <Flex justifyContent={"space-between"} alignItems={"center"}>
        <Text fontWeight={"bold"} fontSize={"xl"}>
          Pesan
        </Text>
        <Button variant={"link"}>
          <SquarePenIcon />
        </Button>
      </Flex>

      <InputGroup>
        <InputLeftElement pointerEvents="none" color={"gray.500"}>
          <SearchIcon />
        </InputLeftElement>
        <Input
          placeholder="Search..."
          bg={"gray.100"}
          rounded={"lg"}
          border={"none"}
          _dark={{ bg: "gray.600" }}
        />
      </InputGroup>

      <Flex flexDir={"column"} gap={4}>
        {users.map((user) => (
          <Flex key={user.id} gap={3} alignItems={"center"}>
            <Avatar
              name={user.full_name}
              src={user.profile_image_url}
              size={"sm"}
            >
              <AvatarBadge boxSize="1em" bg="green.500" />
            </Avatar>
            <Text fontWeight={"semibold"}>{user.full_name}</Text>
          </Flex>
        ))}
      </Flex>
    </Stack>
  );
}
