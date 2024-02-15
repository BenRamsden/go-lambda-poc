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
import { useGetAssetsQuery } from "@/components/gql/generated";

const Home = () => {
  const { data } = useGetAssetsQuery();

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
                <CardContent>Content</CardContent>
              </AccordionContent>
            </AccordionItem>
          </Accordion>
        </Card>

        <Separator className="my-8" />

        <Card className="mt-8">
          <CardHeader>
            <h1 className="font-bold">My Assets</h1>
          </CardHeader>
          <CardContent>
            <ul className="list-none">
              {data &&
                data.assets.map((asset) => {
                  return (
                    <li key={asset.ID} className="mb-4">
                      <h2 className="font-bold">{asset.Name}</h2>
                      <p>{asset.Description}</p>
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

