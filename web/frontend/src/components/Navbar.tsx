import {
  Avatar,
  Box,
  Button,
  Center,
  Flex,
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
import { MoonIcon, PawPrintIcon, SunIcon } from "lucide-react";

export default function Navbar() {
  const { colorMode, toggleColorMode } = useColorMode();
  //   const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Box borderBottom={"4px"} borderBottomColor={"red.400"} px={4}>
        <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
          <Flex alignItems={"center"} gap={2}>
            <PawPrintIcon size={32} />
            <Text fontSize="2xl" fontWeight="bold">
              HaiwanGram
            </Text>
          </Flex>

          <Flex alignItems={"center"}>
            <Stack direction={"row"} spacing={7}>
              <Button onClick={toggleColorMode}>
                {colorMode === "light" ? <MoonIcon /> : <SunIcon />}
              </Button>

              <Menu>
                <MenuButton
                  as={Button}
                  rounded={"full"}
                  variant={"link"}
                  cursor={"pointer"}
                  minW={0}
                >
                  <Avatar
                    name="Ryan Florence"
                    src="https://bit.ly/ryan-florence"
                  />
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
