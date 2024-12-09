import { render, screen } from "@testing-library/react";
import App from "./App";

test("renders all buttons", () => {
  render(<App />);
  const buttonLabels = [
    "Primary",
    "Secondary",
    "Success",
    "Warning",
    "Danger",
    "Info",
    "Light",
    "Dark",
    "Link",
  ];

  buttonLabels.forEach((label) => {
    const button = screen.getByText(label);
    expect(button).toBeInTheDocument();
  });
});
