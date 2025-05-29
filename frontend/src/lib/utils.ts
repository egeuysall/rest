import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: (string | undefined | null | false | Record<string, unknown>)[]) {
  return twMerge(clsx(inputs));
}
