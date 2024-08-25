import { Item } from "@/pages/main";
import { Link } from "react-router-dom";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Badge } from "../ui/badge";
import { Building2, HandCoinsIcon, MapPin } from "lucide-react";
import React from "react";

export default function JobCard({ job }: { job: Item }) {
  return (
    <div className="w-full shadow-sms hover:shadow-lg hover:translate-x-1 transition duration-75">
      <Link to={`/jobs/${job.id}`}>
        <div className="flex border rounded-md border-slate-400 px-3 py-1 w-full">
          <div className="flex items-center mr-1 ">
            <Avatar>
              <AvatarImage
                src="https://sf1-lark-tos.f.mioffice.cn/obj/static-atsx-online-ee-tob/3c46b0f71765aa018256901bcff58378/e7714c27714b83e4c36ef45c69ee49dc6e44d32540880490f6af9a34f47d52f0.png"
                alt="xiaomi"
              />
              <AvatarFallback>小米</AvatarFallback>
            </Avatar>
          </div>
          <div className="flex flex-col w-full p-2">
            <div className="flex justify-between">
              <p className="text-base font-medium">{job.title}</p>
              {/* <div className="flex items-center">
                <JapaneseYen size={14} />
                <p className="text-sky-400 text-pretty">
                  {job.price_down}k - {job.price_up}k
                </p>
              </div> */}
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
            <div className="flex text-xs space-x-4 mb-2">
              <Bad title={job.company_name} logo={<Building2 size={12} />} />
              <Bad
                title={job.city == "" ? "未知" : job.city}
                logo={<MapPin size={12} />}
              />
              <Bad title={job.job_type} />
              <Bad title={job.fulltime ? "全职" : "兼职"} />
            </div>
            <div>
              <p className="text-sm">
                {/* {job.description} */}
                关键词1 关键词2 关键词3
              </p>
            </div>
          </div>
        </div>
      </Link>
    </div>
  );
}

export function Bad({
  title,
  logo,
}: {
  title: string;
  logo?: React.ReactNode;
}) {
  return (
    <Badge
      variant="outline"
      className="flex items-center justify-center space-x-1 font-normal"
    >
      {logo}
      <div>
        <p className="">{title}</p>
      </div>
    </Badge>
  );
}
