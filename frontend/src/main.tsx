import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import MainPage, { mainPageLoader } from "./pages/main";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Job, { jobPageLoader } from "./pages/job";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import Root from "./pages/root";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { ThemeProvider } from "./components/base/theme-provider";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 1000,
    },
  },
});

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        index: true,
        element: <MainPage />,
        loader: mainPageLoader(queryClient),
      },
    ],
  },
  {
    path: "/jobs/:jobId",
    element: <Job />,
    loader: jobPageLoader(queryClient),
  },
]);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider defaultTheme="light" storageKey="lang-them">
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
        <ReactQueryDevtools buttonPosition="bottom-right" />
      </QueryClientProvider>
    </ThemeProvider>
  </StrictMode>
);
