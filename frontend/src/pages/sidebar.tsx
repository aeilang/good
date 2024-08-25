import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { CircleX } from "lucide-react";
import { useRef, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { Search } from "./root";

export default function Sidebar() {
  const [searchParams, setSearchParams] = useSearchParams();
  const companyName = searchParams.get(Search.companyName) || "";
  const query = searchParams.get(Search.search) || "";
  const jobType = searchParams.get(Search.jobType) || "";

  const [cityKey, setCityKey] = useState(+new Date());

  const refCompany = useRef<HTMLInputElement>(null);
  const refSearch = useRef<HTMLInputElement>(null);

  const handleChange = (item: Search, value: string) => {
    if (!value) {
      setSearchParams((prev) => {
        prev.delete(item);
        prev.set(Search.offset, "");
        return prev;
      });
    } else {
      setSearchParams((prev) => {
        prev.set(item, value);
        prev.set(Search.offset, "");
        return prev;
      });
    }
  };

  return (
    <div className="flex">
      <div className="flex items-center justify-center space-x-1 relative">
        {jobType !== "" && (
          <div
            onClick={() => {
              setCityKey(+new Date());
              setSearchParams((prev) => {
                prev.delete(Search.jobType);
                return prev;
              });
            }}
            className="hover:cursor-pointer absolute top-0 right-0 z-10"
          >
            <CircleX
              size={15}
              className="hover:scale-105 text-red-600 transition duration-75 hover:text-red-700"
            />
          </div>
        )}

        <JobTypeSelector
          value={jobType}
          setValue={(str: string) => {
            setSearchParams((prev) => {
              prev.set(Search.jobType, str);
              prev.set(Search.offset, "");
              return prev;
            });
          }}
          key={cityKey}
        />
      </div>
      <div>
        <div className="grid w-full max-w-sm items-center gap-1.5">
          <Label htmlFor="company">公司</Label>
          <div className="flex items-center justify-center">
            <Input
              id="company"
              type="search"
              defaultValue={companyName}
              ref={refCompany}
              className="focus:ring-0"
            />

            <Button
              onClick={() => {
                handleChange(
                  Search.companyName,
                  refCompany.current?.value || ""
                );
              }}
            >
              确定
            </Button>
          </div>
        </div>
      </div>

      <div>
        <div className="grid w-full max-w-sm items-center gap-1.5">
          <Label htmlFor="company">标题</Label>
          <div className="flex items-center justify-center">
            <Input
              id="company"
              type="search"
              defaultValue={query}
              ref={refSearch}
              className="focus:ring-0"
            />

            <Button
              onClick={() => {
                handleChange(Search.search, refSearch.current?.value || "");
              }}
            >
              确定
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}

const JobTypeSelector = ({
  value,
  key,
  setValue,
}: {
  value: string;
  setValue: (value: string) => void;
  key: number;
}) => {
  return (
    <Select defaultValue={value} onValueChange={setValue} key={key}>
      <SelectTrigger className="w-full">
        <SelectValue placeholder="职业类型" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectItem value="校招">校招</SelectItem>
          <SelectItem value="社招">社招</SelectItem>
          <SelectItem value="实习生">实习生</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};
