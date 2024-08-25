import { useState } from "react";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { useSearchParams } from "react-router-dom";
import { Search } from "@/pages/root";
import { CircleX } from "lucide-react";

export default function JobTypeSelect() {
  const [search, setSearch] = useSearchParams();
  const [value, setValue] = useState(() => {
    return search.get(Search.jobType) || "";
  });
  const [key, setKey] = useState(+new Date());

  const handleChange = (v: string) => {
    setValue(v);
    if (!v) {
      setSearch((prev) => {
        prev.delete(Search.offset);
        prev.delete(Search.jobType);
        return prev;
      });
      return;
    }

    setSearch((prev) => {
      prev.delete(Search.offset);
      prev.set(Search.jobType, v);
      return prev;
    });
  };

  return (
    <div className="w-1/5 relative">
      <JobTypeSelector value={value} key={key} setValue={handleChange} />
      {value && (
        <div className="absolute right-1/3 top-1/2 -translate-y-1/2">
          <button
            onClick={() => {
              handleChange("");
              setKey(+new Date());
            }}
          >
            <CircleX size={12} className="text-red-500" />
          </button>
        </div>
      )}
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
        <SelectValue placeholder="招聘类型" />
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
