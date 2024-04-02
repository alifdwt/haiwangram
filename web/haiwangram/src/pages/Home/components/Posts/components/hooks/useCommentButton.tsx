import axiosFetch from "@/config/axiosFetch";
import useToast from "@/hooks/useToast";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ChangeEvent, useCallback, useMemo, useState } from "react";
import { useSelector } from "react-redux";

interface CommentFormValues {
  message: string;
  photo_id: number;
}

export default function useCommentButton({ postId }: { postId: number }) {
  const { accessToken } = useSelector((state: RootState) => state.user);
  const [comment, setComment] = useState("");
  const queryClient = useQueryClient();
  const toast = useToast();
  const form = useMemo(() => {
    return new FormData();
  }, []);

  const mutation = useMutation({
    mutationFn: (newComment: CommentFormValues) =>
      axiosFetch.post("/comments/", newComment, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts", postId] });
      queryClient.invalidateQueries({ queryKey: ["comments", postId] });
      toast("Success", "Berhasil menambahkan komentar", "success");
    },
    onError: () => {
      toast("Error", "Gagal menambahkan komentar", "error");
    },
  });

  const handleCommentChange = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      setComment(e.target.value);
    },
    []
  );

  const handleSubmit = useCallback(() => {
    form.append("message", comment);
    mutation.mutate({ message: comment, photo_id: postId });
    setComment("");
  }, [comment, mutation, postId, form]);

  return { comment, handleCommentChange, handleSubmit, mutation };
}
