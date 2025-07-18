<template>
  <v-card class="bg-v-theme-surface" data-test="devices-list-chooser">
    <DataTable
      v-model:page="page"
      v-model:itemsPerPage="itemsPerPage"
      :headers
      :items="devices"
      :totalCount="numberDevices"
      :loading
      data-test="devices-dataTable"
    >
      <template v-slot:rows>
        <tr v-for="(item, i) in devices" :key="i">
          <td class="pa-0 text-center">
            <v-checkbox
              v-if="props.isSelectable"
              v-model="selected"
              class="mt-5 ml-5"
              density="compact"
              :value="item.uid"
            />
          </td>
          <td class="text-center">
            <router-link
              :to="{ name: 'DeviceDetails', params: { identifier: item.uid } }"
            >
              {{ item.name }}
            </router-link>
          </td>

          <td class="text-center" v-if="item.info">
            <DeviceIcon
              :icon="item.info.id"
              data-test="deviceIcon-component"
            />
            {{ item.info.pretty_name }}
          </td>

          <td class="text-center">
            <v-chip>
              <span
                class="hover-text"
              >
                {{ address(item) }}
              </span>
            </v-chip>
          </td>
        </tr>
      </template>
    </DataTable>
  </v-card>
</template>

<script setup lang="ts">
import { ref, computed, watch } from "vue";
import axios, { AxiosError } from "axios";
import { useStore } from "@/store";
import DataTable from "../DataTable.vue";
import DeviceIcon from "./DeviceIcon.vue";
import handleError from "@/utils/handleError";
import { IDevice } from "@/interfaces/IDevice";
import useSnackbar from "@/helpers/snackbar";

const props = defineProps(["isSelectable"]);
const snackbar = useSnackbar();
const store = useStore();

const headers = [
  {
    text: "",
    value: "selected",
  },
  {
    text: "Hostname",
    value: "hostname",
  },
  {
    text: "Operating System",
    value: "info.pretty_name",
  },
  {
    text: "SSHID",
    value: "namespace",
  },
];

const loading = ref(false);

const filter = ref("");

const itemsPerPage = ref(5);

const page = ref(1);

const devices = computed(
  () => store.getters["devices/getDevicesForUserToChoose"],
);

const numberDevices = computed(
  () => store.getters["devices/getNumberForUserToChoose"],
);

const selected = computed({
  get() {
    return store.getters["devices/getDevicesSelected"];
  },
  set(value) {
    store.commit("devices/setDevicesSelected", value);
  },
});

const getDevices = async (perPageValue: number, pageValue: number) => {
  try {
    loading.value = true;

    const hasDevices = await store.dispatch(
      "devices/setDevicesForUserToChoose",
      {
        perPage: perPageValue,
        page: pageValue,
        filter: filter.value,
        sortStatusField: store.getters["devices/getSortStatusField"],
        sortStatusString: store.getters["devices/getSortStatusString"],
      },
    );

    if (!hasDevices) {
      page.value--;
    }

    loading.value = false;
  } catch (error: unknown) {
    const axiosError = error as AxiosError;
    switch (axios.isAxiosError(error)) {
      case axiosError.response?.status === 403: {
        snackbar.showError("You do not have permission to access this resource.");
        break;
      }
      default: {
        snackbar.showError("An error occurred while fetching devices.");
        break;
      }
    }
    handleError(error);
  }
};

watch([page, itemsPerPage], async () => {
  await getDevices(itemsPerPage.value, page.value);
});

watch(selected, (newValue, oldValue) => {
  if (newValue.length > 3) {
    selected.value = oldValue;
  }
});

const address = (item: IDevice) => `${item.namespace}.${item.name}@${window.location.hostname}`;
</script>
