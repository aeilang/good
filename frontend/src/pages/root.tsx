import { Link, Outlet } from "react-router-dom";
import SearchBar from "@/components/base/searchbar";
import { Avatar } from "@/components/ui/avatar";
import { AvatarFallback, AvatarImage } from "@radix-ui/react-avatar";
import JobTypeSelect from "@/components/base/job-type";
import CitySelect from "@/components/base/city";
import PriceSelect from "@/components/base/price";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { ChevronsDown, ChevronsUp } from "lucide-react";
import { ModeToggle } from "@/components/base/toggle";

export enum Search {
  search = "search",
  priceDown = "price_down",
  priceUp = "price_up",
  companyName = "company_name",
  city = "city",
  jobType = "job_type",
  offset = "offset",
  limit = "limit",
}

export default function Root() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="bg-background container">
      <div className="flex flex-col relative gap-5">
        <div className="sticky z-10 top-0 left-0 right-0 bg-background flex flex-col gap-2 pb-2 pt-2 border-b-2">
          <div className="flex items-center justify-between">
            <div>
              <Link to={"/"}>
                <Avatar>
                  <AvatarImage src="https://lucide.dev/logo.svg" alt="lucide" />
                  <AvatarFallback>IM</AvatarFallback>
                </Avatar>
              </Link>
            </div>
            <div className="w-1/2 flex items-center gap-1">
              <SearchBar />
              <Button
                variant={"link"}
                className="text-sky-500"
                onClick={() => {
                  setIsOpen((prev) => !prev);
                }}
              >
                过滤
                {isOpen ? <ChevronsUp size={15} /> : <ChevronsDown size={15} />}
              </Button>
            </div>
            <div>
              <ModeToggle />
            </div>
          </div>
          {isOpen && (
            <div className="flex space-x-2 items-center justify-between">
              <JobTypeSelect />
              <CitySelect item={Search.city} holder="城市" />
              <CitySelect item={Search.companyName} holder="公司名称" />
              <PriceSelect />
            </div>
          )}
        </div>

        <div>
          <Outlet />
        </div>
      </div>
    </div>
  );
}
