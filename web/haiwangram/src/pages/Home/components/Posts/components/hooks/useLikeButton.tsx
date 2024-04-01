import axiosFetch from "@/config/axiosFetch";
import useToast from "@/hooks/useToast";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { MouseEvent } from "react";
import { useSelector } from "react-redux";

interface LikeFormValues {
  photo_id: number;
}

export default function useLikeButton({
  isLiked,
  postId,
}: {
  isLiked: boolean;
  postId: number;
}) {
  const queryClient = useQueryClient();
  const toast = useToast();
  const { accessToken } = useSelector((state: RootState) => state.user);

  const mutation = useMutation({
    mutationFn: (newLike: LikeFormValues) => {
      return axiosFetch.post("/likes/", newLike, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts"] });
      toast("Success", "Berhasil menyukai post", "success");
    },
    onError: () => {
      toast("Error", "Gagal menyukai post", "error");
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (newLike: LikeFormValues) => {
      return axiosFetch.delete("/likes/", {
        data: newLike,
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts"] });
      toast("Success", "Suka post ini berhasil dibatalkan", "success");
    },
    onError: () => {
      toast("Error", "Gagal menyukai post", "error");
    },
  });

  const handleLike = (e: MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();

    if (!isLiked) {
      mutation.mutate({ photo_id: postId });
      return;
    }

    deleteMutation.mutate({ photo_id: postId });
  };

  return {
    handleLike,
  };
}
