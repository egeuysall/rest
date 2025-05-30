"use client";

import { useEffect } from "react";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    console.error("Error in payload page:", error);
  }, [error]);

  return (
    <div className="min-h-screen bg-background flex items-center justify-center">
      <div className="w-full max-w-3xl mx-auto px-4 py-12">
        <div className="bg-white dark:bg-neutral-900 rounded-lg shadow-lg border border-neutral-200 dark:border-neutral-800 p-8">
          <h2 className="text-2xl font-bold text-neutral-900 dark:text-neutral-100 mb-4">
            Something went wrong!
          </h2>
          <p className="text-neutral-600 dark:text-neutral-400 mb-6">
            {error.message ||
              "An unexpected error occurred while loading the payload."}
          </p>
          <button
            onClick={reset}
            className="px-4 py-2 bg-neutral-800 dark:bg-neutral-700 text-white rounded-md hover:bg-neutral-700 dark:hover:bg-neutral-600 transition-colors"
          >
            Try again
          </button>
        </div>
      </div>
    </div>
  );
}
