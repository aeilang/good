import { Bad } from "@/components/base/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import {
  QueryClient,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import {
  AlignHorizontalSpaceAround,
  ArrowLeft,
  ArrowLeftCircle,
  Building2,
  Flower2,
  HandCoinsIcon,
  MapPin,
  Moon,
  PartyPopper,
  SendHorizonal,
  Tally1,
  Tally2,
} from "lucide-react";
import {
  Link,
  LoaderFunctionArgs,
  useLoaderData,
  useNavigate,
} from "react-router-dom";

export const jobQuery = (id: string) =>
  queryOptions({
    queryKey: ["job", id],
    queryFn: async () => {
      try {
        const data = await fetch("http://localhost:8080/api/v1/post/" + id);
        return data.json();
      } catch {
        return null;
      }
    },
  });

export type Job = {
  href: string;
  company_name: string;
  title: string;
  city: string;
  fulltime: boolean;
  job_type: string;
  description: string;
  requirement: string;
  price_down: number;
  price_up: number;
};

export const jobPageLoader =
  (queryClient: QueryClient) =>
  async ({ params }: LoaderFunctionArgs) => {
    const id = params["jobId"] || "";

    await queryClient.ensureQueryData(jobQuery(id));

    return { id };
  };

export default function Job() {
  const { id } = useLoaderData() as Awaited<
    ReturnType<ReturnType<typeof jobPageLoader>>
  >;

  const navagate = useNavigate();

  const { data } = useSuspenseQuery(jobQuery(id));

  const job = data as Job;

  return (
    <div className="flex bg-background items-center justify-center container h-screen w-screen">
      <div className="flex flex-col space-y-2">
        <div>
          <Button
            variant={"link"}
            onClick={() => {
              navagate(-1);
            }}
            className="text-sky-400"
          >
            <div className="flex items-center justify-center">
              <ArrowLeft />
              <p className="text-lg">返回</p>
            </div>
          </Button>
        </div>
        <div className="flex items-center justify-between">
          <div className="flex items-center justify-center gap-2">
            <PartyPopper />
            <h1 className="text-2xl">{job.title}</h1>
          </div>
          <Badge
            variant={"secondary"}
            className="flex items-center justify-center space-x-2 text-base"
          >
            <HandCoinsIcon className="text-yellow-400" />
            <p className="font-semibold text-sky-400">
              {job.price_down} - {job.price_up} K
            </p>
          </Badge>
        </div>
        <div className="flex space-x-2 text-sm">
          <Bad title={job.company_name} logo={<Building2 size={12} />} />
          <Bad
            title={job.city == "" ? "未知" : job.city}
            logo={<MapPin size={12} />}
          />
          <Bad title={job.job_type} />
          <Bad title={job.fulltime ? "全职" : "兼职"} />
        </div>

        <Separator />
        <div className="flex flex-col gap-2">
          <div className="flex items-center">
            <Tally1 />
            <p className="text-lg">岗位职责:</p>
          </div>
          <div style={{ whiteSpace: "pre-wrap" }}>
            <p>{job.description}</p>
          </div>
        </div>

        <Separator />
        <div className="flex flex-col gap-2">
          <div className="flex items-center">
            <Tally2 />
            <p className="text-lg">任职要求:</p>
          </div>
          <div style={{ whiteSpace: "pre-wrap" }}>
            <p>{job.requirement}</p>
          </div>
        </div>
        <div className="flex pt-6">
          <Link to={job.href} target="_blank" className="mx-auto">
            <Button variant={"destructive"}>
              <div className="flex items-center justify-center gap-2">
                <p className="text-lg">立即投递</p>
                <SendHorizonal />
              </div>
            </Button>
          </Link>
        </div>
      </div>
    </div>
  );
}
