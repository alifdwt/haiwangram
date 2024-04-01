import axiosFetch from "@/config/axiosFetch";
import useToast from "@/hooks/useToast";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useSelector } from "react-redux";

export default function useCreatePost() {
  const queryClient = useQueryClient();
  const { user, accessToken } = useSelector((state: RootState) => state.user);
  const [title, setTitle] = useState("");
  const [caption, setCaption] = useState("");
  const [image, setImage] = useState<File | null>(null);

  const toast = useToast();
  const form = new FormData();

  const mutation = useMutation({
    mutationFn: (newPost: FormData) =>
      axiosFetch.post("/photos/", newPost, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    onSuccess: () => {
      toast("Success", "Postingan anda berhasil diproses", "success");
      setTimeout(() => {
        queryClient.invalidateQueries({ queryKey: ["posts"] });
      }, 3000);
    },
    onError: () => {
      toast("Gagal", "Postingan anda gagal diproses", "error");
    },
  });

  const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };

  const handleCaptionChange = (
    event: React.ChangeEvent<HTMLTextAreaElement>
  ) => {
    setCaption(event.target.value);
  };

  const handleSubmit = () => {
    form.append("title", title);
    form.append("caption", caption);
    form.append("image", image as File);

    mutation.mutate(form);
    setTitle("");
    setCaption("");
    setImage(null);
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setImage(file);
    }
  };

  const handleFileUpload = () => {
    const input = document.createElement("input");
    input.type = "file";
    input.id = "image";
    input.name = "image";

    //@ts-expect-error next-line
    input.onchange = handleFileChange;
    input.click();
  };

  return {
    title,
    caption,
    image,
    setImage,
    user,
    handleTitleChange,
    handleCaptionChange,
    handleSubmit,
    handleFileUpload,
  };
}
