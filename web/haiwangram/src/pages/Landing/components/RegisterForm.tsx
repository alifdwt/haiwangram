import {
  AbsoluteCenter,
  Box,
  Button,
  Center,
  Divider,
  Flex,
  Heading,
  Stack,
  Text,
} from "@chakra-ui/react";
import { FacebookIcon } from "lucide-react";
import Register from "./actions/Register";
import Login from "./actions/Login";

export default function RegisterForm() {
  return (
    <Flex h={"full"} flexDir={"column"} justifyContent={"space-between"} p={4}>
      <Center flexDir={"column"} gap={2}>
        <Heading fontSize={"4xl"} fontWeight={"extrabold"}>
          Mari buat akunmu!
        </Heading>
        <Stack spacing={2} align={"center"} maxW={"md"} w={"full"}>
          <Button
            w={"full"}
            colorScheme="red"
            leftIcon={
              <svg
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="currentColor"
                viewBox="0 0 18 19"
                width={18}
                height={18}
              >
                <path
                  fill-rule="evenodd"
                  d="M8.842 18.083a8.8 8.8 0 0 1-8.65-8.948 8.841 8.841 0 0 1 8.8-8.652h.153a8.464 8.464 0 0 1 5.7 2.257l-2.193 2.038A5.27 5.27 0 0 0 9.09 3.4a5.882 5.882 0 0 0-.2 11.76h.124a5.091 5.091 0 0 0 5.248-4.057L14.3 11H9V8h8.34c.066.543.095 1.09.088 1.636-.086 5.053-3.463 8.449-8.4 8.449l-.186-.002Z"
                  clip-rule="evenodd"
                />
              </svg>
            }
          >
            Daftar dengan Google
          </Button>
          <Button w={"full"} colorScheme="facebook" leftIcon={<FacebookIcon />}>
            Daftar dengan Facebook
          </Button>
        </Stack>
        <Box position="relative" p={4}>
          <Divider borderWidth={"2px"} w={"md"} mx={"auto"} />
          <AbsoluteCenter bg="white" px="4" color="gray.500">
            atau
          </AbsoluteCenter>
        </Box>
        <Register />
        <Text w={"md"} color={"gray"} fontSize={"sm"}>
          Lorem, ipsum dolor sit amet consectetur adipisicing elit. Quidem quam
          molestias soluta, ea ab excepturi veniam perferendis, earum doloribus,
          accusamus nesciunt modi optio iure culpa quae fugiat. Eligendi quos
          sunt amet rerum adipisci, corporis, enim facilis, asperiores expedita
          est aliquid non repellat quas iure. Quibusdam debitis ut praesentium
          nemo delectus?
        </Text>
      </Center>
      <Divider borderWidth={"2px"} w={"md"} mx={"auto"} />
      <Center flexDir={"column"} gap={2}>
        <Text color={"gray"}>Sudah punya akun?</Text>
        <Login />
      </Center>
    </Flex>
  );
}
