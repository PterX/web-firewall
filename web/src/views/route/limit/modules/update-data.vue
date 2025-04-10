<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import { updateForwardLimitPolicy } from '@/service/api';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { checkIpMask } from '@/utils/ip_check';

const { formRef, validate } = useNaiveForm();

const showModal = defineModel<boolean>('show');

const props = defineProps<{
  row: any;
}>();

const emit = defineEmits<{
  (e: 'close'): void;
}>();

const loading = ref(false);

const formValue = ref<any>({
  protocol: '',
  dipAny: true,
  dip: '',
  portType: 'sport',
  port: '',
  sipAny: true,
  sip: '',
  limit: null,
  speed: 'mb/s',

  comment: ''
});

const rules = computed<any>(() => {
  // inside computed to make locale reactive, if not apply i18n, you can define it without computed
  const { defaultRequiredRule } = useFormRules();

  return {
    sip: [
      defaultRequiredRule,
      {
        trigger: ['input', 'change'],
        validator(_rule: any, value: string) {
          const state = value.split(',').every((item: string) => {
            return checkIpMask(item);
          });
          if (!state) {
            return new Error($t('page.firewallPolicy.ipValidationFailure'));
          }
          return true;
        }
      }
    ],

    dip: [
      defaultRequiredRule,
      {
        trigger: ['input', 'change'],
        validator(_rule: any, value: string) {
          const state = value.split(',').every((item: string) => {
            return checkIpMask(item);
          });
          if (!state) {
            return new Error($t('page.firewallPolicy.ipValidationFailure'));
          }
          return true;
        }
      }
    ],
    port: [
      defaultRequiredRule,
      {
        trigger: ['input', 'change'],
        pattern: /^\d[\d,-]*\d$/,
        message: $t('page.firewallPolicy.portValidationFailure')
      },
      {
        trigger: ['input', 'change'],
        validator(_rule: any, value: string) {
          const state = value.split(',').every((item: string) => {
            const intItem = Number.parseInt(item, 10);
            if (intItem < 0 || intItem > 65535) {
              return false;
            }
            return true;
          });
          if (!state) {
            return new Error($t('page.firewallPolicy.portValidationFailure'));
          }
          return true;
        }
      }
    ],

    policy: [defaultRequiredRule]
  };
});

function initData() {
  formValue.value = {
    protocol: 'any',
    dipAny: true,
    dip: '',
    portType: 'sport',
    port: '',
    sipAny: true,
    sip: '',
    limit: null,
    speed: 'mb/s',
    comment: ''
  };
  emit('close');
}

async function onSubmit() {
  await validate();
  //  提交数据
  loading.value = true;

  const { error } = await updateForwardLimitPolicy({
    id: props.row.id,
    ...formValue.value,
    add: !(formValue.value.add === 1 || formValue.value.add === 3),
    position: formValue.value.add > 2 ? formValue.value.position : 0,
    sip: formValue.value.sipAny ? '' : formValue.value.sip,
    dip: formValue.value.dipAny ? '' : formValue.value.dip,
    port: formValue.value.protocol === '' ? '' : formValue.value.port,
    portType: formValue.value.protocol === '' ? '' : formValue.value.portType
  });
  loading.value = false;
  if (error) return;
  window.$message?.success($t('common.modifySuccess'));
  showModal.value = false;
}

async function enterModal() {
  formValue.value = props.row;
  formValue.value.sipAny = formValue.value.sip === '';
  formValue.value.dipAny = formValue.value.dip === '';
  loading.value = false;
}
</script>

<template>
  <NModal
    v-model:show="showModal"
    :mask-closable="false"
    preset="card"
    class="w-700px"
    :title="$t('common.edit')"
    :bordered="false"
    :segmented="{
      content: true
    }"
    @after-leave="initData"
    @after-enter="enterModal"
  >
    <NSpin :show="loading">
      <NForm
        ref="formRef"
        :model="formValue"
        label-width="100px"
        label-placement="left"
        label-align="left"
        :rules="rules"
        class="ml-20px mr-30px"
      >
        <NFormItem :label="$t('page.firewallPolicy.protocol')" path="protocol">
          <!-- <NInput v-model:value="formValue.protocol" /> -->
          <NSelect
            v-model:value="formValue.protocol"
            :options="[
              {
                label: $t('page.firewallPolicy.all'),
                value: ''
              },
              {
                label: 'tcp',
                value: 'tcp'
              },
              {
                label: 'udp',
                value: 'udp'
              }
            ]"
          />
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.sourceIp')" path="sipAny">
          <NRadioGroup v-model:value="formValue.sipAny" name="radiogroup">
            <NSpace>
              <NRadio :value="true">
                {{ $t('page.firewallPolicy.allIp') }}
              </NRadio>
              <NRadio :value="false">
                {{ $t('page.firewallPolicy.partialIp') }}
              </NRadio>
            </NSpace>
          </NRadioGroup>
        </NFormItem>

        <NFormItem v-if="!formValue.sipAny" label=" " path="sip">
          <NSpace vertical :size="14" class="w-full">
            <NInput
              v-model:value="formValue.sip"
              type="textarea"
              :autosize="{
                minRows: 1,
                maxRows: 5
              }"
            />
            <span class="mb-30px mt-10px font-size-14px text-truegray-400">
              {{ $t('page.firewallPolicy.ipTip') }}
            </span>
          </NSpace>
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.destIp')" path="dipAny">
          <NRadioGroup v-model:value="formValue.dipAny" name="radiogroup">
            <NSpace>
              <NRadio :value="true">
                {{ $t('page.firewallPolicy.allIp') }}
              </NRadio>
              <NRadio :value="false">
                {{ $t('page.firewallPolicy.partialIp') }}
              </NRadio>
            </NSpace>
          </NRadioGroup>
        </NFormItem>

        <NFormItem v-if="!formValue.dipAny" label=" " path="dip">
          <NSpace vertical :size="14" class="w-full">
            <NInput
              v-model:value="formValue.dip"
              type="textarea"
              :autosize="{
                minRows: 1,
                maxRows: 5
              }"
            />
            <span class="mb-30px mt-10px font-size-14px text-truegray-400">
              {{ $t('page.firewallPolicy.ipTip') }}
            </span>
          </NSpace>
        </NFormItem>

        <div v-if="formValue.protocol !== ''">
          <NFormItem :label="$t('page.firewallPolicy.port')" path="portType">
            <NRadioGroup v-model:value="formValue.portType" name="radiogroup">
              <NSpace>
                <NRadio value="sport">
                  {{ $t('page.firewallPolicy.sourcePort') }}
                </NRadio>
                <NRadio value="dport">
                  {{ $t('page.firewallPolicy.destPort') }}
                </NRadio>
              </NSpace>
            </NRadioGroup>
          </NFormItem>

          <NFormItem label=" " path="port">
            <NSpace vertical :size="14" class="w-full">
              <NInput v-model:value="formValue.port" />
              <span class="mb-30px mt-10px font-size-14px text-truegray-400">
                {{ $t('page.firewallPolicy.portTip') }}
              </span>
            </NSpace>
          </NFormItem>
        </div>

        <NFormItem :label="$t('page.firewallPolicy.speed')" path="limit">
          <!-- <NInput v-model:value="formValue.policy" /> -->
          <NSpace>
            <NInputNumber v-model:value="formValue.limit" />

            <NSelect
              v-model:value="formValue.speed"
              class="w-215px"
              :options="[
                {
                  label: 'kb/s',
                  value: 'kb/s'
                },
                {
                  label: 'mb/s',
                  value: 'mb/s'
                },
                {
                  label: 'kb/m',
                  value: 'kb/m'
                },
                {
                  label: 'mb/m',
                  value: 'mb/m'
                }
              ]"
            />
          </NSpace>
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.comment')" path="comment">
          <NInput v-model:value="formValue.comment" type="textarea" />
        </NFormItem>
      </NForm>
    </NSpin>
    <template #footer>
      <NSpace justify="end">
        <NButton @click="showModal = false">{{ $t('common.cancel') }}</NButton>
        <NButton v-throttle="onSubmit" type="primary">{{ $t('common.confirm') }}</NButton>
      </NSpace>
    </template>
  </NModal>
</template>
