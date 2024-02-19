import Token from "@/components/auth/token";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import Header from "@/components/ui/header";
import { toast } from "sonner";

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Separator } from "@/components/ui/separator";
import {
  GetAssetsDocument,
  useCreateAssetMutation,
  useGetAssetsQuery,
  useGetPanicLazyQuery,
} from "@/components/gql/generated";
import { Button } from "@/components/ui/button";
import { RefreshCw } from "lucide-react";
import { AssetView } from "@/components/assets/asset-view";
import { Input } from "@/components/ui/input";
import { useForm } from "react-hook-form";

const Extras = () => {
  const [getPanic] = useGetPanicLazyQuery({
    fetchPolicy: "network-only",
    onError: () =>
      toast("Query Paniced", {
        description: "Sentry will capture this error",
        action: {
          label: "View",
          onClick: () => {
            open(
              "https://jugo-digital-ltd.sentry.io/issues/?environment=sandbox&project=4506772933509120"
            );
          },
        },
      }),
  });

  return (
    <div className="flex flex-row">
      <Button
        className="mr-4"
        variant={"destructive"}
        onClick={() => {
          getPanic({
            variables: {
              message: "This is a panic message",
            },
          });
        }}
      >
        Trigger Sentry Error
      </Button>
    </div>
  );
};

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
      refetchQueries: [
        {
          query: GetAssetsDocument,
        },
      ],
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
  const { data: assetsData, refetch: refetchAssets } = useGetAssetsQuery({
    fetchPolicy: "network-only",
  });

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
            <AccordionItem value="item-extra">
              <CardHeader>
                <AccordionTrigger>
                  <h1 className="font-bold">Extras</h1>
                </AccordionTrigger>
              </CardHeader>
              <AccordionContent>
                <CardContent>
                  <Extras />
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
              <Button variant="outline" onClick={() => refetchAssets()}>
                <RefreshCw />
              </Button>
            </div>
          </CardHeader>
          <CardContent>
            <ul className="list-none">
              {assetsData &&
                assetsData.assets.map((asset, index) => {
                  return (
                    <li key={asset.ID} className="mb-4">
                      <AssetView asset={asset} />
                      {index !== assetsData.assets.length - 1 && (
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

