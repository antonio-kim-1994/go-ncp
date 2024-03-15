-- 테이블 생성 SQL - ncloud_contract_summary_list
CREATE TABLE ncloud_contract_summary_list
(
    `id`                  INT            NOT NULL    AUTO_INCREMENT COMMENT 'id',
    `contract_type_code`  VARCHAR(50)    NOT NULL    COMMENT '계약 유형 코드. getContractSummaryListResponse.contractSummaryList[].contractType.code',
    `contract_code_name`  VARCHAR(50)    NOT NULL    COMMENT '계약 코드 이름. getContractSummaryListResponse.contractSummaryList[].contractType.codeName',
    `contract_region`     VARCHAR(50)    NOT NULL    COMMENT '계약 리전. getContractSummaryListResponse.contractSummaryList[].regionCode',
    `contract_count`      INT            NOT NULL    COMMENT '계약 수. getContractSummaryListResponse.contractSummaryList[].contractCount',
    `write_date`          DATE           NOT NULL    COMMENT '작성 일자. API 집계 일자',
    PRIMARY KEY (id)
);

-- 테이블 Comment 설정 SQL - ncloud_contract_summary_list
ALTER TABLE ncloud_contract_summary_list COMMENT '네이버 클라우드 계약 현황 리스트';

-- 테이블 생성 SQL - ncloud_product_cost_list
CREATE TABLE ncloud_product_cost_list
(
    `id`                        INT            NOT NULL    AUTO_INCREMENT COMMENT 'id',
    `product_demand_type_code`  VARCHAR(30)    NOT NULL    COMMENT '상품 청구 유형 코드. getProductDemandCostListResponse.productDemandCostList[].productDemandType.code',
    `product_demand_code_name`  VARCHAR(50)    NOT NULL    COMMENT '상품 청구 코드 이름. getProductDemandCostListResponse.productDemandCostList[].productDemandType.codeName',
    `product_use_amount`        INT            NOT NULL    COMMENT '상품 이용 금액. getProductDemandCostListResponse.productDemandCostList[].useAmount [상품 과금 집계 금액]',
    `product_demand_amount`     INT            NOT NULL    COMMENT '상품 청구 금액. getProductDemandCostListResponse.productDemandCostList[].demandAmount [상품 과금 최종 결정 금액]',
    `write_date`                DATE           NOT NULL    COMMENT '작성 일자. API 집계 일자',
    PRIMARY KEY (id)
);

-- 테이블 Comment 설정 SQL - ncloud_product_cost_list
ALTER TABLE ncloud_product_cost_list COMMENT '네이버 클라우드 상품 비용 리스트';