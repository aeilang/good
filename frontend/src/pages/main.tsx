import JobCard from "@/components/base/card";
import {
  QueryClient,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import { LoaderFunctionArgs, useLoaderData } from "react-router-dom";

import Pagina from "@/components/base/pagination";

export const jobsQuery = (input: string) =>
  queryOptions({
    queryKey: ["jobs", input],
    queryFn: async () => {
      try {
        const data = await fetch("http://localhost:8080/api/v1/posts" + input);
        return data.json();
      } catch {
        return null;
      }
    },
  });

export type Item = {
  city: string;
  company_name: string;
  description: string;
  fulltime: boolean;
  id: string;
  job_type: string;
  price_down: number;
  price_up: number;
  title: string;
  total_count: number;
};

export const mainPageLoader =
  (queryClient: QueryClient) =>
  async ({ request }: LoaderFunctionArgs) => {
    const url = new URL(request.url);
    const search = url.search;

    await queryClient.ensureQueryData(jobsQuery(search));
    return { search };
  };

export default function MainPage() {
  const { search } = useLoaderData() as Awaited<
    ReturnType<ReturnType<typeof mainPageLoader>>
  >;

  const { data } = useSuspenseQuery(jobsQuery(search));

  const jobs = data as Item[];

  const count = jobs ? jobs[0].total_count : 0;

  return (
    <div className="relative">
      <div className="flex flex-col space-y-4 w-full mx-auto">
        {jobs && jobs.map((job) => <JobCard job={job} />)}
        {!jobs && <NoData />}
      </div>

      <div className="pb-5 sticky z-10 bg-background w-full bottom-0">
        <Pagina count={count} />
      </div>
    </div>
  );
}

const NoData = () => {
  return (
    <div className="h-screen mx-auto flex items-center">
      <p>没有更多数据了</p>
    </div>
  );
};
