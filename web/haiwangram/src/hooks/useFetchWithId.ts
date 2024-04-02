import axiosFetch from "@/config/axiosFetch";
import { useQuery } from "@tanstack/react-query";

interface IUserFetchWithId {
  id: string;
  queryKey: string;
  fetchRoute: string;
}

export default function useFetchWithId({
  id,
  queryKey,
  fetchRoute,
}: IUserFetchWithId) {
  const { data, isLoading } = useQuery({
    queryKey: [queryKey, id],
    queryFn: async () => {
      const { data } = await axiosFetch.get(`${fetchRoute}/${id}`);
      return data;
    },
  });

  return { data, isLoading };
}
