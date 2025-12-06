import React from "react";
import {
  BarChart as ReBarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
  LabelList,
} from "recharts";
import type { ChartItem } from "../api/report";

interface BarChartProps {
  data: ChartItem[];
}

export default function BarChart({ data }: BarChartProps) {
  return (
    <div style={{ width: "100%", height: 300, marginBottom: 30 }}>
      <ResponsiveContainer>
        <ReBarChart data={data} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="area" />
          <YAxis domain={[0, 110]} unit="%" />
          <Tooltip />
          <Legend />
          <Bar dataKey="value" fill="#8884d8" name="Nilai">
            <LabelList dataKey="value" position="top" formatter={(val) => `${val.toFixed(1)}%`} />
          </Bar>
        </ReBarChart>
      </ResponsiveContainer>
    </div>
  );
}
