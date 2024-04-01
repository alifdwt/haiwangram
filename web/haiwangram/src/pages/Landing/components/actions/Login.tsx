import {
  Box,
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Input,
  InputGroup,
  InputRightElement,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  useDisclosure,
} from "@chakra-ui/react";
import useLogin from "./hooks/useLogin";
import { EyeIcon, EyeOffIcon } from "lucide-react";

export default function Login() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const {
    handleSubmit,
    register,
    errors,
    onSubmit,
    mutation,
    showPassword,
    handleShowPassword,
  } = useLogin();

  return (
    <>
      <Button
        onClick={onOpen}
        variant={"outline"}
        _hover={{ bg: "primary.700", color: "white" }}
        w={"md"}
      >
        Masuk
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Masuk</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Box as={"form"} onSubmit={handleSubmit(onSubmit)}>
              <Flex direction={"column"} gap={2}>
                <FormControl isInvalid={Boolean(errors.email)}>
                  <Input
                    id="email"
                    type="email"
                    placeholder="Email"
                    {...register("email", {
                      required: "Email is required!",
                      pattern: {
                        value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
                        message: "invalid email address!",
                      },
                    })}
                  />
                  <FormErrorMessage>
                    {errors.email && errors.email.message?.toString()}
                  </FormErrorMessage>
                </FormControl>
                <FormControl isInvalid={Boolean(errors.password)}>
                  <InputGroup>
                    <Input
                      id="password"
                      type={showPassword ? "text" : "password"}
                      placeholder="Password"
                      {...register("password", {
                        required: "Password is required!",
                        minLength: {
                          value: 6,
                          message: "Password must be at least 6 characters!",
                        },
                      })}
                    />
                    <InputRightElement>
                      <Button size="sm" onClick={handleShowPassword}>
                        {showPassword ? <EyeIcon /> : <EyeOffIcon />}
                      </Button>
                    </InputRightElement>
                  </InputGroup>
                  <FormErrorMessage>
                    {errors.password && errors.password.message?.toString()}
                  </FormErrorMessage>
                </FormControl>
                <Button
                  type="submit"
                  bg={"primary.700"}
                  color={"white"}
                  _hover={{ bg: "primary.700", color: "white" }}
                  isLoading={mutation.isPending}
                  loadingText={"Memproses"}
                >
                  Masuk
                </Button>
              </Flex>
            </Box>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="blue" mr={3} onClick={onClose}>
              Close
            </Button>
            <Button variant="ghost">Secondary Action</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}
