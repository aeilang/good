import { HandCoins, MoveRight } from "lucide-react";
import CitySelect from "./city";
import { Search } from "@/pages/root";

export default function PriceSelect() {
  return (
    <div className="flex w-2/5 items-center justify-between gap-2">
      <div>
        <HandCoins size={20} />
      </div>
      <div className="flex items-center justify-center gap-2">
        <CitySelect item={Search.priceDown} holder="下限" />
        <p>K</p>
      </div>
      <div>
        <MoveRight size={12} />
      </div>
      <div className="flex items-center justify-center gap-2">
        <CitySelect item={Search.priceUp} holder="上限" />
        <p>K</p>
      </div>
    </div>
  );
}
