import { useRef } from "react";
import { Input } from "../ui/input";
import { Check, CircleX } from "lucide-react";
import { useSearchParams } from "react-router-dom";
import { Search } from "@/pages/root";

export default function CitySelect({
  item,
  holder,
}: {
  item: Search;
  holder: string;
}) {
  const ref = useRef<HTMLInputElement>(null);
  const [search, setSearch] = useSearchParams();
  const defaultValue = search.get(item) || "";

  const handleClick = () => {
    const value = ref.current?.value;

    if (!value) {
      setSearch((prev) => {
        prev.delete(item);
        prev.delete(Search.offset);
        return prev;
      });
    } else {
      setSearch((prev) => {
        prev.set(item, value);
        prev.delete(Search.offset);
        return prev;
      });
    }
  };

  return (
    <div className="relative">
      <Input
        placeholder={holder}
        ref={ref}
        defaultValue={defaultValue}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            handleClick();
          }
        }}
      />

      <div className="absolute right-2 top-1/2 -translate-y-1/2">
        <div className="flex gap-2">
          {ref.current?.value && (
            <div>
              <button
                onClick={() => {
                  if (ref.current) {
                    ref.current.value = "";
                    handleClick();
                  }
                }}
              >
                <CircleX size={14} className="text-red-500" />
              </button>
            </div>
          )}
          {!ref.current?.value && (
            <div>
              <button onClick={handleClick}>
                <Check
                  size={14}
                  className="text-sky-600 hover:scale-105 hover:text-sky-400"
                />
              </button>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
