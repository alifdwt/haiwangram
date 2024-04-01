import axiosFetch from "@/config/axiosFetch";
import useToast from "@/hooks/useToast";
import User from "@/interface/User";
// import { login } from "@/slices/user/userSlice";
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
  Spinner,
  useDisclosure,
} from "@chakra-ui/react";
import { useMutation } from "@tanstack/react-query";
import { EyeIcon, EyeOffIcon } from "lucide-react";
import { useState } from "react";
import { useForm } from "react-hook-form";
// import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";

interface RegisterFormValues {
  email: string;
  username: string;
  full_name: string;
  birth_date: string;
  password: string;
  description: string;
}

export default function Register() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  // const dispatch = useDispatch()
  const {
    handleSubmit,
    register,
    reset,
    formState: { errors },
  } = useForm<RegisterFormValues>();
  const navigate = useNavigate();
  const toast = useToast();
  const [showPassword, setShowPassword] = useState(false);

  const mutation = useMutation({
    mutationFn: (newUser: User) => {
      return axiosFetch.post("/auth/register", newUser);
    },
    onSuccess: () => {
      toast("Success", "Akun berhasil dibuat", "success");
      // dispatch(login(data.data.token))
      reset();
      navigate("/");
    },
    onError: () => {
      toast("Error", "Akun gagal dibuat", "error");
    },
  });

  const onSubmit = (data: RegisterFormValues) => {
    mutation.mutate(data);
  };

  const handleConfimPassword = () => {
    setShowPassword(!showPassword);
  };

  return (
    <>
      <Button
        onClick={onOpen}
        bg={"primary.700"}
        color={"white"}
        _hover={{ bg: "primary.700" }}
        w={"md"}
      >
        Buat akun
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Buat akun</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Box as="form" onSubmit={handleSubmit(onSubmit)}>
              <Flex direction={"column"} gap={4}>
                <FormControl isInvalid={Boolean(errors.email)}>
                  <Input
                    id="email"
                    type="email"
                    placeholder="Masukkan email"
                    {...register("email", {
                      required: "Email harus diisi",
                      pattern: {
                        value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
                        message: "Email tidak valid",
                      },
                    })}
                  />
                  <FormErrorMessage>
                    {errors.email && errors.email.message?.toString()}
                  </FormErrorMessage>
                </FormControl>
                <FormControl isInvalid={Boolean(errors.username)}>
                  <Input
                    id="username"
                    type="text"
                    placeholder="Masukkan username"
                    {...register("username", {
                      required: "Username harus diisi",
                    })}
                  />
                  <FormErrorMessage>
                    {errors.username && errors.username.message?.toString()}
                  </FormErrorMessage>
                </FormControl>
                <FormControl isInvalid={Boolean(errors.full_name)}>
                  <Input
                    id="full_name"
                    type="text"
                    placeholder="Masukkan nama lengkap"
                    {...register("full_name", {
                      required: "Nama lengkap harus diisi",
                    })}
                  />
                  <FormErrorMessage>
                    {errors.full_name && errors.full_name.message?.toString()}
                  </FormErrorMessage>
                </FormControl>
                <FormControl isInvalid={Boolean(errors.birth_date)}>
                  <Input
                    id="birth_date"
                    type="date"
                    placeholder="Masukkan tanggal lahir"
                    {...register("birth_date", {
                      required: "Tanggal lahir harus diisi",
                    })}
                  />
                  <FormErrorMessage>
                    {errors.birth_date && errors.birth_date.message?.toString()}
                  </FormErrorMessage>
                </FormControl>
                <FormControl isInvalid={Boolean(errors.password)}>
                  <InputGroup>
                    <Input
                      id="password"
                      type={showPassword ? "text" : "password"}
                      placeholder="Masukkan password"
                      {...register("password", {
                        required: "Password harus diisi",
                        minLength: {
                          value: 6,
                          message: "Password minimal 6 karakter",
                        },
                      })}
                    />
                    <InputRightElement>
                      <Button
                        size={"sm"}
                        px={2}
                        bg={"transparent"}
                        _hover={{ bg: "transparent" }}
                        onClick={handleConfimPassword}
                      >
                        {showPassword ? <EyeIcon /> : <EyeOffIcon />}
                      </Button>
                    </InputRightElement>
                  </InputGroup>
                  <FormErrorMessage>
                    {errors.password && errors.password.message?.toString()}
                  </FormErrorMessage>
                </FormControl>
                <FormControl isInvalid={Boolean(errors.description)}>
                  <Input
                    id="description"
                    type="text"
                    placeholder="Masukkan deskripsi"
                    {...register("description")}
                  />
                  <FormErrorMessage>
                    {errors.description &&
                      errors.description.message?.toString()}
                  </FormErrorMessage>
                </FormControl>
                <Button
                  type="submit"
                  bg={"primary.700"}
                  color={"white"}
                  _hover={{ bg: "primary.700" }}
                  w={"full"}
                >
                  {mutation.isPending ? <Spinner /> : "Buat akun"}
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
