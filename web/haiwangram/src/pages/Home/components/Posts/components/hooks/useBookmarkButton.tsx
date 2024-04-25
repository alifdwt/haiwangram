import axiosFetch from "@/config/axiosFetch";
import useToast from "@/hooks/useToast";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { MouseEvent } from "react";
import { useSelector } from "react-redux";

interface BookmarkFormValues {
  photo_id: number;
}

export default function useBookmarkButton({
  isBookmarked,
  postId,
}: {
  isBookmarked: boolean;
  postId: number;
}) {
  const queryClient = useQueryClient();
  const toast = useToast();
  const { accessToken } = useSelector((state: RootState) => state.user);

  const mutation = useMutation({
    mutationFn: (newBookmark: BookmarkFormValues) => {
      return axiosFetch.post("/bookmarks/", newBookmark, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts"] });
      toast("Success", "Berhasil menambahkan bookmark", "success");
    },
    onError: () => {
      toast("Error", "Gagal menambahkan bookmark", "error");
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (newBookmark: BookmarkFormValues) => {
      return axiosFetch.delete("/bookmarks/", {
        data: newBookmark,
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts"] });
      queryClient.invalidateQueries({ queryKey: ["post", postId] });
      toast("Success", "Berhasil menghapus bookmark", "success");
    },
    onError: () => {
      toast("Error", "Gagal menghapus bookmark", "error");
    },
  });

  const handleBookmark = (e: MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();

    if (!isBookmarked) {
      mutation.mutate({ photo_id: postId });
      return;
    }

    deleteMutation.mutate({ photo_id: postId });
  };

  return {
    handleBookmark,
  };
}
