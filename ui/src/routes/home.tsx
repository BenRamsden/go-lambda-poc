import Token from "@/components/auth/token";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import Header from "@/components/ui/header";

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Separator } from "@/components/ui/separator";
import {
  Asset,
  useCreateAssetMutation,
  useGetAssetsLazyQuery,
} from "@/components/gql/generated";
import { Button } from "@/components/ui/button";
import { RefreshCw } from "lucide-react";
import { useCallback, useEffect, useState } from "react";
import { AssetView } from "@/components/assets/asset-view";
import { Input } from "@/components/ui/input";
import { useForm } from "react-hook-form";

const CreateAsset = () => {
  const [createAsset] = useCreateAssetMutation();
  const { register, handleSubmit } = useForm({
    defaultValues: {
      Name: "",
      Description: "",
    },
  });

  const onSubmit = async (data: { Name: string; Description: string }) => {
    await createAsset({
      variables: {
        input: {
          Name: data.Name,
          Description: data.Description,
          URI: "https://www.google.com",
        },
      },
      refetchQueries: ["GetAssets"],
    });
  };

  return (
    <form className="container my-4" onSubmit={handleSubmit(onSubmit)}>
      <Input className="mb-4" placeholder="Name" {...register("Name")} />
      <Input
        className="mb-4"
        placeholder="Description"
        {...register("Description")}
      />
      <Button type="submit">Create</Button>
    </form>
  );
};

const Home = () => {
  const [localAssets, setLocalAssets] = useState<Asset[]>([]);

  const [fetchAssets, { data, loading: loadingAssets }] = useGetAssetsLazyQuery(
    {
      fetchPolicy: "network-only",
    }
  );

  useEffect(() => {
    fetchAssets();
  }, []);

  useEffect(() => {
    if (!loadingAssets && data?.assets)
      setLocalAssets(
        [...data.assets].sort(
          (a, b) =>
            new Date(b.CreatedAt).getTime() - new Date(a.CreatedAt).getTime()
        )
      );
  }, [loadingAssets]);

  const handleRefreshAssets = useCallback(() => {
    setLocalAssets([]);
    fetchAssets();
  }, [fetchAssets]);

  return (
    <div>
      <Header />

      <main className="container">
        <Card className="mt-8">
          <Accordion type="single" collapsible>
            <AccordionItem value="item-token">
              <CardHeader>
                <AccordionTrigger>
                  <h1 className="font-bold">Token</h1>
                </AccordionTrigger>
              </CardHeader>
              <AccordionContent>
                <CardContent>
                  <Token />
                </CardContent>
              </AccordionContent>
            </AccordionItem>
          </Accordion>
        </Card>

        <Card className="mt-8">
          <Accordion type="single" collapsible>
            <AccordionItem value="item-create-asset">
              <CardHeader>
                <AccordionTrigger>
                  <h1 className="font-bold">New Asset</h1>
                </AccordionTrigger>
              </CardHeader>
              <AccordionContent>
                <CardContent>
                  <CreateAsset />
                </CardContent>
              </AccordionContent>
            </AccordionItem>
          </Accordion>
        </Card>

        <Separator className="my-8" />

        <Card className="mt-8">
          <CardHeader>
            <div className="flex flex-row justify-between items-center">
              <h1 className="font-bold">My Assets</h1>
              <Button variant="outline" onClick={() => handleRefreshAssets()}>
                <RefreshCw />
              </Button>
            </div>
          </CardHeader>
          <CardContent>
            <ul className="list-none">
              {!loadingAssets &&
                localAssets.map((asset, index) => {
                  return (
                    <li key={asset.ID} className="mb-4">
                      <AssetView asset={asset} />
                      {index !== localAssets.length - 1 && (
                        <Separator className="my-4" />
                      )}
                    </li>
                  );
                })}
            </ul>
          </CardContent>
        </Card>
      </main>
    </div>
  );
};

export default Home;

