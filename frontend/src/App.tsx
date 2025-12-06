import React, { useState, useEffect } from "react";
import FilterForm from "./components/FilterForm";
import BarChart from "./components/BarChart";
import DataTable from "./components/DataTable";
import type { ChartItem, TableItem } from "./api/report";
import { getReports } from "./api/report";

function App() {
  const [chartData, setChartData] = useState<ChartItem[]>([]);
  const [tableData, setTableData] = useState<TableItem[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function handleView(area_id?: string, dateFrom?: string, dateTo?: string) {
    setLoading(true);
    setError(null);

    try {
      const data = await getReports({ area_id, dateFrom, dateTo });

      setChartData(data.chart ?? []);
      setTableData(data.table ?? []);
    } catch (err) {
      console.error(err);
      setError("Gagal mengambil data laporan");
      setChartData([]);
      setTableData([]);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    handleView();
  }, []);

  return (
    <div
      style={{
        height: "100%",
        minHeight: "100vh",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        padding: "20px",
        backgroundColor: "#f0f2f5",
        boxSizing: "border-box",
      }}
    >
      {/* Filter Form */}
      <div style={{ width: "100%", maxWidth: "900px", marginBottom: "20px" }}>
        <FilterForm onSubmit={handleView} />
      </div>

      {/* Loading */}
      {loading && (
        <div
          style={{
            padding: "20px",
            textAlign: "center",
            color: "#555",
            fontSize: "18px",
          }}
        >
          Loading data...
        </div>
      )}

      {/* Error */}
      {error && (
        <div
          style={{
            padding: "20px",
            marginBottom: "20px",
            width: "100%",
            maxWidth: "900px",
            textAlign: "center",
            color: "#fff",
            backgroundColor: "#e74c3c",
            borderRadius: "5px",
          }}
        >
          {error}
        </div>
      )}

      {/* Chart */}
      {!loading && !error && chartData?.length > 0 && (
        <div
          style={{
            width: "100%",
            maxWidth: "900px",
            backgroundColor: "#fff",
            borderRadius: "10px",
            padding: "20px",
            marginBottom: "30px",
            boxShadow: "0 4px 12px rgba(0,0,0,0.1)",
          }}
        >
          <BarChart data={chartData} />
        </div>
      )}

      {/* Table */}
      {!loading && !error && tableData?.length > 0 && (
        <div
          style={{
            width: "100%",
            maxWidth: "900px",
            backgroundColor: "#fff",
            borderRadius: "10px",
            padding: "20px",
            boxShadow: "0 4px 12px rgba(0,0,0,0.1)",
          }}
        >
          <DataTable data={tableData} />
        </div>
      )}

      {/* Message */}
      {!loading && !error && chartData?.length === 0 && tableData?.length === 0 && (
        <div style={{ padding: "20px", color: "#555", fontSize: "16px" }}>
          Tidak ada data untuk ditampilkan
        </div>
      )}
    </div>
  );
}

export default App;
