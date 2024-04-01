import axiosFetch from "@/config/axiosFetch";
import useToast from "@/hooks/useToast";
import { login } from "@/slices/user/userSlice";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";

interface LoginFormValues {
  email: string;
  password: string;
}

export default function useLogin() {
  const [showPassword, setShowPassword] = useState(false);
  const {
    handleSubmit,
    register,
    reset,
    formState: { errors },
  } = useForm<LoginFormValues>();
  const toast = useToast();
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const mutation = useMutation({
    mutationFn: (user: LoginFormValues) => {
      return axiosFetch.post("/auth/login", user);
    },
  });

  const onSubmit = (data: LoginFormValues) => {
    mutation.mutate(data, {
      onSuccess: (data) => {
        toast("Success", "Login berhasil", "success");
        dispatch(login(data.data.token));
        reset();
        navigate("/");
      },
      onError: () => {
        toast("Error", "Login gagal", "error");
      },
    });
  };

  const handleShowPassword = () => {
    setShowPassword((prevState) => !prevState);
  };

  return {
    handleSubmit,
    register,
    errors,
    onSubmit,
    mutation,
    showPassword,
    handleShowPassword,
  };
}
