import {
  Avatar,
  Box,
  Button,
  Center,
  Flex,
  Input,
  InputGroup,
  InputLeftElement,
  Link,
  Menu,
  MenuButton,
  MenuDivider,
  MenuItem,
  MenuList,
  Stack,
  Text,
  useColorMode,
  //   useDisclosure,
} from "@chakra-ui/react";
import {
  ArrowDown,
  BellIcon,
  MoonIcon,
  PawPrintIcon,
  SearchIcon,
  SunIcon,
} from "lucide-react";
import { LogOutButton } from "./components/LogButton";
import { useSelector } from "react-redux";
import { RootState } from "@/store";

export default function Topbar() {
  const { colorMode, toggleColorMode } = useColorMode();
  const { user } = useSelector((state: RootState) => state.user);

  return (
    <>
      <Box px={4} bg={"white"} _dark={{ bg: "gray.700" }}>
        <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
          <Link
            href="/"
            display={"flex"}
            alignItems={"center"}
            gap={2}
            _hover={{ textDecoration: "none" }}
          >
            <Box bg={"primary.700"} borderRadius={"lg"} color={"white"} p={2}>
              <PawPrintIcon />
            </Box>
            <Text fontSize="2xl" fontWeight="bold">
              HaiwanGram
            </Text>
          </Link>

          <InputGroup maxW={"450px"}>
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

          <Flex alignItems={"center"}>
            <Stack direction={"row"} spacing={4}>
              <Button>
                <BellIcon />
              </Button>
              <Button onClick={toggleColorMode}>
                {colorMode === "light" ? <MoonIcon /> : <SunIcon />}
              </Button>

              <Menu>
                <MenuButton
                  as={Button}
                  rounded={"full"}
                  variant={"ghost"}
                  cursor={"pointer"}
                  minW={0}
                >
                  <Flex gap={2} alignItems={"center"}>
                    <Avatar
                      name={user?.full_name}
                      src={user?.profile_image_url}
                      size={"sm"}
                    />
                    <Text>
                      {user?.full_name?.split(" ")[0]}{" "}
                      {
                        user?.full_name?.split(" ")[
                          user.full_name?.split(" ").length - 1
                        ]
                      }
                    </Text>
                    <ArrowDown />
                  </Flex>
                </MenuButton>
                <MenuList alignItems={"center"}>
                  <br />
                  <Center>
                    <Avatar
                      name={user?.full_name}
                      src={user?.profile_image_url}
                      size="xl"
                      cursor="pointer"
                    />
                  </Center>
                  <br />
                  <Center>
                    <p>{user?.full_name}</p>
                  </Center>
                  <br />
                  <MenuDivider />
                  <MenuItem>Karyamu</MenuItem>
                  <MenuItem>Pengaturan Akun</MenuItem>
                  <MenuItem>
                    <LogOutButton />
                  </MenuItem>
                </MenuList>
              </Menu>
            </Stack>
          </Flex>
        </Flex>
      </Box>
    </>
  );
}
