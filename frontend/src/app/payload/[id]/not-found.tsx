import Link from "next/link";

export default function NotFound() {
  return (
    <div className="min-h-screen bg-background flex items-center justify-center">
      <div className="w-full max-w-3xl mx-auto px-4 py-12">
        <div className="bg-white dark:bg-neutral-900 rounded-lg shadow-lg border border-neutral-200 dark:border-neutral-800 p-8">
          <h2 className="text-2xl font-bold text-neutral-900 dark:text-neutral-100 mb-4">
            Payload Not Found
          </h2>
          <p className="text-neutral-600 dark:text-neutral-400 mb-6">
            The payload you're looking for doesn't exist or has expired.
          </p>
          <Link
            href="/"
            className="px-4 py-2 bg-neutral-800 dark:bg-neutral-700 text-white rounded-md hover:bg-neutral-700 dark:hover:bg-neutral-600 transition-colors inline-block"
          >
            Go back home
          </Link>
        </div>
      </div>
    </div>
  );
}
