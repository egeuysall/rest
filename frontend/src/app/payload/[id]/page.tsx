import { notFound } from "next/navigation";
import CodeBlock from "@/components/code-block";
import { Clock, RefreshCw } from "lucide-react";

type PageProps = {
  params: Promise<{ id: string }>;
};

export const dynamic = "force-dynamic";
export const revalidate = 0;

type PayloadData = {
  [key: string]: string | number | boolean | null | PayloadData | PayloadData[];
};

type ApiResponse = {
  data: PayloadData;
  expires_at: string;
  remaining_reads: number;
};

export default async function Page({ params }: PageProps) {
  const resolvedParams = await params;
  const id = resolvedParams.id;
  if (!id) {
    console.error("No ID provided");
    notFound();
  }

  const apiKey = process.env.NEXT_PUBLIC_REST_API_KEY;
  if (!apiKey) {
    console.error("NEXT_PUBLIC_REST_API_KEY is missing");
    throw new Error("NEXT_PUBLIC_REST_API_KEY is missing");
  }

  console.log("Fetching payload for ID:", id);
  const res = await fetch(`https://restapi.egeuysal.com/v1/payload/${id}`, {
    headers: {
      Authorization: `Bearer ${apiKey}`,
    },
    cache: "no-store",
  });

  console.log("Response status:", res.status);
  if (res.status === 404) {
    console.error("Payload not found");
    notFound();
  }

  if (!res.ok) {
    const errorText = await res.text();
    console.error("API error response:", errorText);
    throw new Error(`API error: ${res.status} - ${errorText}`);
  }

  const payload: ApiResponse = await res.json();
  console.log("Successfully fetched payload:", payload);

  if (!payload || !payload.data) {
    console.error("Invalid payload format:", payload);
    throw new Error("Invalid payload format");
  }

  return (
    <main className="min-h-screen bg-background flex items-center justify-center">
      <div className="w-full max-w-3xl mx-auto px-4 py-12">
        <div className="bg-white dark:bg-neutral-900 rounded-lg shadow-lg border border-neutral-200 dark:border-neutral-800 overflow-hidden">
          <div className="p-8 border-b border-neutral-200 dark:border-neutral-800">
            <h1 className="text-3xl font-bold text-foreground text-center mb-6">
              Payload Details
            </h1>
            <div className="flex justify-center gap-8">
              <div className="flex items-center gap-2 text-sm text-neutral-600 dark:text-neutral-400">
                <RefreshCw className="w-4 h-4" />
                <span>Remaining Reads: {payload.remaining_reads}</span>
              </div>
              <div className="flex items-center gap-2 text-sm text-neutral-600 dark:text-neutral-400">
                <Clock className="w-4 h-4" />
                <span>
                  Expires: {new Date(payload.expires_at).toLocaleString()}
                </span>
              </div>
            </div>
          </div>
          <div className="p-8">
            <CodeBlock
              code={JSON.stringify(payload.data, null, 2)}
              language="json"
              fileName="payload"
            />
          </div>
        </div>
      </div>
    </main>
  );
}
