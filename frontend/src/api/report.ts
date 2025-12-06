export interface ChartItem {
  area: string;
  value: number;
}

export interface TableItem {
  brand: string;
  area: string;
  value: number;
}

export interface ReportResponse {
  chart: ChartItem[];
  table: TableItem[];
}

export interface ReportParams {
  dateFrom?: string;
  dateTo?: string;
  area_id?: string;
}

export async function getReports({
  dateFrom,
  dateTo,
  area_id,
}: ReportParams): Promise<ReportResponse> {
  const params = new URLSearchParams();

  if (dateFrom) params.append("dateFrom", dateFrom);
  if (dateTo) params.append("dateTo", dateTo);
  if (area_id) params.append("area_id", area_id);

  const res = await fetch(`http://localhost:8080/api/reports?${params.toString()}`, {
    headers: { "X-API-Key": "TEST123456" },
  });

  if (!res.ok) {
    throw new Error("Failed to fetch reports");
  }

  return res.json() as Promise<ReportResponse>;
}

export interface Area {
  id: string;
  name: string;
}

export async function getAreas(): Promise<Area[]> {
  const res = await fetch("http://localhost:8080/api/areas", {
    headers: { "X-API-Key": "TEST123456" },
  });

  if (!res.ok) {
    throw new Error("Failed to fetch areas");
  }

  return res.json() as Promise<Area[]>;
}
