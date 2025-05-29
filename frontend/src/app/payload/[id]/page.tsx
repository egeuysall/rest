"use client";

import { useEffect, useState } from "react";
import { use } from "react";
import { notFound } from "next/navigation";
import CodeBlock from "@/components/code-block";
import { Clock, RefreshCw } from "lucide-react";

type PayloadData = {
  [key: string]:
    | string
    | number
    | boolean
    | null
    | PayloadData
    | Array<PayloadData>;
};

type ApiResponse = {
  data: PayloadData;
  expires_at: string;
  remaining_reads: number;
};

type Props = {
  params: Promise<{ id: string }>;
};

export default function PayloadPage({ params }: Props) {
  const { id } = use(params);
  const [payload, setPayload] = useState<ApiResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPayload = async () => {
      try {
        const apiKey = process.env.NEXT_PUBLIC_REST_API_KEY;
        if (!apiKey) {
          throw new Error("API key is not configured");
        }

        const apiUrl = `https://restapi.egeuysal.com/v1/payload/${id}`;
        console.log("Fetching from:", apiUrl);

        const res = await fetch(apiUrl, {
          headers: {
            Authorization: `Bearer ${apiKey}`,
          },
        });

        if (!res.ok) {
          if (res.status === 404) {
            notFound();
          }
          if (res.status === 401) {
            throw new Error(
              "Authentication failed - please check API key configuration"
            );
          }
          throw new Error(`API request failed with status ${res.status}`);
        }

        const data = await res.json();
        setPayload(data);
      } catch (err) {
        console.error("Error fetching payload:", err);
        setError(
          err instanceof Error ? err.message : "Failed to fetch payload"
        );
      }
    };

    fetchPayload();
  }, [id]);

  if (error) {
    return (
      <div className="min-h-screen bg-background flex items-center justify-center">
        <div className="text-red-500">Error: {error}</div>
      </div>
    );
  }

  if (!payload) {
    return (
      <div className="min-h-screen bg-background flex items-center justify-center">
        <div className="text-neutral-600">Loading...</div>
      </div>
    );
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
