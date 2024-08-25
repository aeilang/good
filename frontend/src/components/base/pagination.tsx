import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import { Search } from "@/pages/root";
import { useSearchParams } from "react-router-dom";

export default function Pagina({ count }: { count: number }) {
  const [search, setSearch] = useSearchParams();
  const str1 = search.get(Search.offset);
  const offset = str1 ? Number(str1) : 0;

  const str2 = search.get(Search.limit);
  const limit = str2 ? Number(str2) : 10;
  const page = offset / limit + 1;
  const maxPage = Math.ceil(count / limit);

  const handleClick = (offest: number) => {
    search.set(Search.offset, String(offest));
    setSearch(search);
  };

  if (count == 0) {
    return;
  }

  return (
    <div>
      <div className="mx-auto min-w-full">
        <Pagination>
          <PaginationContent>
            {page > 1 && (
              <PaginationItem>
                <PaginationPrevious
                  onClick={() => handleClick(offset - limit)}
                />
              </PaginationItem>
            )}

            {Array.from(
              { length: Math.min(Math.floor(maxPage / 2), 3) },
              (_, i) => {
                const length = Math.min(Math.floor(maxPage / 2), 3);
                if (page - (length - 1 - i) > 0 && i !== length - 1) {
                  return (
                    <PaginationItem>
                      <PaginationLink
                        onClick={() =>
                          handleClick((page - (length - 1 - i) - 1) * limit)
                        }
                      >
                        {page - (length - 1 - i)}
                      </PaginationLink>
                    </PaginationItem>
                  );
                }
              }
            )}

            {Array.from(
              { length: Math.min(Math.floor(maxPage / 2), 5) },
              (_, i) => {
                if (page + i <= maxPage) {
                  return (
                    <PaginationItem>
                      <PaginationLink
                        onClick={() => handleClick((page + i - 1) * limit)}
                        isActive={i === 0}
                      >
                        {page + i}
                      </PaginationLink>
                    </PaginationItem>
                  );
                }
              }
            )}

            {page !== maxPage && (
              <PaginationItem>
                <PaginationLink
                  onClick={() => handleClick((maxPage - 1) * limit)}
                >
                  {maxPage}
                </PaginationLink>
              </PaginationItem>
            )}

            {page < maxPage && (
              <PaginationItem>
                <PaginationNext onClick={() => handleClick(offset + limit)} />
              </PaginationItem>
            )}
          </PaginationContent>
        </Pagination>
      </div>
    </div>
  );
}
