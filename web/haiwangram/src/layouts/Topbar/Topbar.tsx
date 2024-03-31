import {
  Avatar,
  Box,
  Button,
  Center,
  Flex,
  Input,
  InputGroup,
  InputLeftElement,
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

export default function Topbar() {
  const { colorMode, toggleColorMode } = useColorMode();
  //   const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Box px={4} bg={"white"} _dark={{ bg: "gray.700" }}>
        <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
          <Flex alignItems={"center"} gap={2}>
            <Box bg={"primary"} borderRadius={"lg"} color={"white"} p={2}>
              <PawPrintIcon />
            </Box>
            <Text fontSize="2xl" fontWeight="bold">
              HaiwanGram
            </Text>
          </Flex>

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
                      name="Ryan Florence"
                      src="https://bit.ly/ryan-florence"
                      size={"sm"}
                    />
                    <Text>Ryan Florence</Text>
                    <ArrowDown />
                  </Flex>
                </MenuButton>
                <MenuList alignItems={"center"}>
                  <br />
                  <Center>
                    <Avatar
                      name="Ryan Florence"
                      src="https://bit.ly/ryan-florence"
                      size="xl"
                      cursor="pointer"
                    />
                  </Center>
                  <br />
                  <Center>
                    <p>Ryan Florence</p>
                  </Center>
                  <br />
                  <MenuDivider />
                  <MenuItem>Karyamu</MenuItem>
                  <MenuItem>Pengaturan Akun</MenuItem>
                  <MenuItem>Keluar</MenuItem>
                </MenuList>
              </Menu>
            </Stack>
          </Flex>
        </Flex>
      </Box>
    </>
  );
}
