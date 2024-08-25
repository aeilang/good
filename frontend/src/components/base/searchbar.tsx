import { CircleX, Search as SearchIcon } from "lucide-react";
import { Input } from "../ui/input";
import { useSearchParams } from "react-router-dom";
import { Search } from "@/pages/root";
import { useRef } from "react";

export default function SearchBar() {
  const [search, setSearch] = useSearchParams();
  const query = search.get(Search.search) || "";
  const ref = useRef<HTMLInputElement>(null);

  const handleSubmit = () => {
    const value = ref.current?.value;
    if (!value) {
      setSearch((prev) => {
        prev.delete(Search.offset);
        prev.delete(Search.search);
        return prev;
      });
      return;
    }

    setSearch((prev) => {
      prev.delete(Search.offset);
      prev.set(Search.search, value);
      return prev;
    });
  };

  return (
    <div className="w-full">
      <div className="relative">
        <Input
          className="rounded-full"
          placeholder="python..."
          defaultValue={query}
          name="query"
          ref={ref}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleSubmit();
            }
          }}
        />
        <div className="absolute -translate-y-1/2 top-1/2 right-4">
          <div className="flex gap-5">
            {ref.current?.value && (
              <button
                onClick={() => {
                  if (ref.current) {
                    ref.current.value = "";
                    handleSubmit();
                  }
                }}
              >
                <CircleX size={12} className="text-red-600" />
              </button>
            )}
            <button onClick={handleSubmit}>
              <SearchIcon size={20} />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
