# Copyright 2019 Volvo Cars
#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# ”License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#  http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# “AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

defmodule Lin.Master.App do

  use GenServer;
  require Logger;

  defstruct [
    # channels: nil,
    signal_base_pid: nil
  ]

  #Client

  def starthelper do
    start_link(:hejsan, SignalBase.Application.make_signal_broker_name "lin")
  end

  def start_link(name, signal_base_pid) do
    GenServer.start_link(__MODULE__, {signal_base_pid}, name: name)
  end

  #def read(pid, channel_names),
  #  do: GenServer.call(pid, {:read, channel_names})

  #Server 

  def init({signal_base_pid}) do    
    SignalBase.register_listeners(signal_base_pid, ["CCMLIN18Fr03"], self(), self())
    # Lin.Scheduler.run_pattern(:test_lin_schedule, file, "CcmLin18ScheduleTable2", 1)
    # :timer.sleep(500)
    Logger.debug "STARTED"
    {:ok, %__MODULE__{signal_base_pid: signal_base_pid}}
  end

  #Server
  def handle_cast({:signal, msg}, state) do
    msg.name_values
    |> Enum.map(fn touple = {name, value} ->
      case {name, value} do
        {"CCMLIN18Fr03", :arbitration} ->
           Logger.debug "need to send some stuff #{inspect touple}"
          send_our_data(state.signal_base_pid)
        _ ->
          Logger.debug "irrelant #{inspect touple}"
      end
    end)
    {:noreply, state}
  end

  def send_our_data(signal_base_pid) do
    SignalBase.publish(signal_base_pid, [{"FuHeatrAmbAirP", 13}, {"FuHeatrDiagRqrd", 22}], self())
  end


  # #Server
  # def init(signal_base_pid, channel_names, file) do    
  #   SignalBase.register_listeners(signal_base_pid, channel_names, self(), self())
  #   Lin.Scheduler.run_pattern(:test_lin_schedule, file, "CcmLin18ScheduleTable2", 1)
  #   :timer.sleep(500)
  # end

  # def handle_call({:read, channel_names}, _from, state) do
  #   read_values = channel_names |>
  #   Enum.map(fn(channel) ->
  #     case :ets.lookup(state.channels, channel) do
  #       [{channel, value}] -> {channel, value}
  #       [] -> {channel, :empty}
  #     end
  #   end)

  #   {:reply, cached_values, state}
  # end
end
