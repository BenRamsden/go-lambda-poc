import { Asset } from "../gql/generated";
import { formatDistance } from "date-fns";

type AssetViewProps = {
  asset: Asset;
};

const AssetView = ({ asset }: AssetViewProps) => {
  return (
    <div className="flex flex-row justify-between items-center">
      <div className="flex flex-col">
        <h1>{asset.Name}</h1>
        <p className="font-light">{asset.Description}</p>
      </div>
      <p className="font-extralight">
        {formatDistance(asset.CreatedAt, new Date())}
      </p>
    </div>
  );
};

export { AssetView };

