<!--suppress ALL -->
<template>
    <b-table :striped="striped" :narrowed="narrowed" :bordered="bordered" class="common-table"
                    :data="datas" @page-change="(page) => {currentPage=page;$emit('pageChange', page)}" 
                    :total="total" :backend-pagination="paginated" :paginated = "paginated"
                    :per-page="perPage" :detailed="detailed" :detail-key="detailKey"
                    :default-sort="defaultSort" default-sort-direction="asc" :current-page.sync="currentPage">
                    <template slot-scope="props">
                        <b-table-column v-for="(item, index) in columnProps" :key="`table-column-${index}`"
                        :field="item.field" :label="item.label" :width="item.width" :sortable="item.sortable">
                            <template v-if="item.customSlot">
                              <slot :row="props.row" :name="`column-${item.field}-${props.index}`"/>
                            </template>
                            <span v-else>{{ item.isNumbering ? (props.index + 1 + (perPage * (currentPage - 1))) 
                              : props.row[item.field] }}</span>
                        </b-table-column>
                    </template>

                    <template v-if="isDetail" slot="detail" slot-scope="props">
                      <slot name="table-detail" :props="props"/>
                    </template>

                    <template slot="empty">
                        <empty-states/>
                    </template>
    </b-table>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "annotation";
import { isEmpty } from "lodash";

import TableColumn from "../models/TableColumn";

@Component
export default class CommonTable extends Vue {
  // paging
  @Prop({ default: 0 })
  private perPage: number;

  @Prop({ default: 0 })
  private total: number;

  // styling
  @Prop({ default: true })
  private narrowed: boolean;

  @Prop({ default: true })
  private bordered: boolean;

  @Prop({ default: true })
  private striped: boolean;

  @Prop({ default: "" })
  private detailKey: string;

  // sorting
  @Prop({ default: "" })
  private defaultSort: string;

  // detail
  @Prop({ default: false })
  private isDetail: boolean;

  // required
  @Prop({ default: [], required: true })
  private columnProps: Array<TableColumn>;

  @Prop({ default: [], required: true })
  private datas: Array<any>;

  private currentPage: number = 1;

  private get paginated() {
    return this.perPage > 0 || this.total > 0;
  }

  private get detailed() {
    return !isEmpty(this.detailKey);
  }
}
</script>